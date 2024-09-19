// package main
package utils

import (
	"context"
	"log/slog"
	"math/rand"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/users"
)

const (
	SHUFFLE_NUM = 100
)

var (
	dsn       string
	db_client *ent.Client
	ctx       context.Context
)

/* 参加情報用 */
type RegistrationInfo struct {
	user_id     int
	rank        int
	record_id   int
	session_bit int
}

func newRegiStration(info ent.EVENT_RECORDS) RegistrationInfo {
	var ret RegistrationInfo
	ret.user_id = info.UserID
	ret.rank = getRankByUserID(info.UserID)
	ret.record_id = info.ID
	ret.session_bit = info.ParticipatesBit
	return ret
}

func getRankByUserID(user_id int) int {
	// user_idからrankを取得
	user, err := db_client.USERS.Query().
		Where(
			users.IDEQ(user_id),
		).
		First(ctx)
	if err != nil {
		slog.Error("Failed to get rank of %d: %v", user_id, err)
		return -1
	}
	return user.Rank
}

/* マッチング結果用 */
type Pair struct {
	First  RegistrationInfo
	Second RegistrationInfo
}

func newPair(f, s RegistrationInfo) Pair {
	var ret Pair
	ret.First = f
	ret.Second = s
	return ret
}

/* マッチングに使う関数群 */
func Matching(event_id int) {
	var i int
	// DBの用意
	dsn := config.GetDSN()
	db_client, err := ent.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open connection to database: %v", err)
		return
	}
	defer db_client.Close()
	ctx := context.Background()

	var matching_list []Pair
	matching_list = []Pair{}
	/* セッション１のマッチング */
	// EVENT_RECORDSからセッション１に参加するユーザを取得する
	info, err := db_client.EVENT_RECORDS.Query().
		Where(
			event_records.EventIDEQ(event_id),
			event_records.Or(
				event_records.ParticipatesBitEQ(1),
				event_records.ParticipatesBitEQ(3),
				event_records.ParticipatesBitEQ(5),
				event_records.ParticipatesBitEQ(7),
			),
		).
		All(ctx)
	if err != nil {
		logErrorExceptNotFound(err, "Failed to get matching info of session 1")
	} else {
		// 各レベルごとにシャッフル
		shuffled := shuffleRegistration(info)
		// 先頭から2人ずつをマッチングさせる
		matching_list = makeMatchPair(shuffled)
		// マッチング結果をSESSIONSテーブルに格納
		for i = 0; i < len(matching_list); i++ {
			db_client.SESSIONS.Create().
				SetUserID(matching_list[i].First.user_id).
				SetMatchedUserID(matching_list[i].Second.user_id).
				SetRecordID(matching_list[i].First.record_id).
				SetStatus("MATCHED").
				Save(ctx)
		}
	}

	/* セッション２のマッチング */
	// セッション１のマッチング結果から，セッション２に参加しないものを削除
	for i = 0; i < len(matching_list); i++ {
		if matching_list[i].First.session_bit == 1 || matching_list[i].First.session_bit == 5 {
			matching_list[i].First.user_id = -1
		}
		if matching_list[i].Second.session_bit == 1 || matching_list[i].Second.session_bit == 5 {
			matching_list[i].Second.user_id = -1
		}
	}
	i = 0
	for i < len(matching_list) { // 両ユーザが削除されたマッチングをリストから消す
		if matching_list[i].First.user_id == -1 && matching_list[i].Second.user_id == -1 {
			matching_list = append(matching_list[:i], matching_list[i+1:]...)
		} else {
			i++
		}
	}
	// セッション２でマッチングに加わるユーザを取得する
	info2, err := db_client.EVENT_RECORDS.Query().
		Where(
			event_records.EventIDEQ(event_id),
			event_records.Or(
				event_records.ParticipatesBitEQ(2),
				event_records.ParticipatesBitEQ(6),
			),
		).
		All(ctx)
	if err != nil {
		logErrorExceptNotFound(err, "Failed to get matching info of session 2")
		// matching_listをそのまま使う
		matching_list = remakeList(matching_list, []RegistrationInfo{})
	} else {
		// 新たに加える方はシャッフル
		shuffled := shuffleRegistration(info2)
		// matching_listと合体
		matching_list = remakeList(matching_list, shuffled)
	}
	// 連続するマッチング同士で，Firstを入れ替える
	matching_list = session2_swapping(matching_list)
	// DBに格納
	// マッチング結果をSESSIONSテーブルに格納
	for i := 0; i < len(matching_list); i++ {
		db_client.SESSIONS.Create().
			SetUserID(matching_list[i].First.user_id).
			SetMatchedUserID(matching_list[i].Second.user_id).
			SetRecordID(matching_list[i].First.record_id).
			SetStatus("MATCHED").
			Save(ctx)
	}

	/* セッション３のマッチング */
	// セッション２のマッチング結果から，セッション３に参加しないものを削除
	for i = 0; i < len(matching_list); i++ {
		if matching_list[i].First.session_bit == 2 || matching_list[i].First.session_bit == 3 {
			matching_list[i].First.user_id = -1
		}
		if matching_list[i].Second.session_bit == 2 || matching_list[i].Second.session_bit == 3 {
			matching_list[i].Second.user_id = -1
		}
	}
	i = 0
	for i < len(matching_list) { // 両ユーザが削除されたマッチングをリストから消す
		if matching_list[i].First.user_id == -1 && matching_list[i].Second.user_id == -1 {
			matching_list = append(matching_list[:i], matching_list[i+1:]...)
		} else {
			i++
		}
	}
	// セッション２でマッチングに加わるユーザを取得する
	info3, err := db_client.EVENT_RECORDS.Query().
		Where(
			event_records.EventIDEQ(event_id),
			event_records.Or(
				event_records.ParticipatesBitEQ(4),
				event_records.ParticipatesBitEQ(5),
			),
		).
		All(ctx)
	if err != nil {
		logErrorExceptNotFound(err, "Failed to get matching info of session 3")
		// matching_listをそのまま使う
		matching_list = remakeList(matching_list, []RegistrationInfo{})
	} else {
		// 新たに加える方はシャッフル
		shuffled := shuffleRegistration(info3)
		// matching_listと合体
		matching_list = remakeList(matching_list, shuffled)
	}
	// 連続するマッチング同士で，Firstを入れ替える
	matching_list = session3_swapping(matching_list)
	// マッチング結果をSESSIONSテーブルに格納
	for i := 0; i < len(matching_list); i++ {
		db_client.SESSIONS.Create().
			SetUserID(matching_list[i].First.user_id).
			SetMatchedUserID(matching_list[i].Second.user_id).
			SetRecordID(matching_list[i].First.record_id).
			SetStatus("MATCHED").
			Save(ctx)
	}
}

func logErrorExceptNotFound(err error, contextMessage string) {
	if !ent.IsNotFound(err) {
		slog.Error("%s: %v", contextMessage, err)
	}
}

func shuffleRegistration(info []*ent.EVENT_RECORDS) []RegistrationInfo {
	// rankごとに配列を分ける
	rank_pool := make([][]RegistrationInfo, 5)
	for _, user_info := range info {
		add := newRegiStration(*user_info)
		rank_pool[add.rank-1] = append(rank_pool[add.rank-1], add)
	}
	// 各rankでシャッフルを実施
	for rank := 0; rank < 5; rank++ {
		if len(rank_pool[rank]) > 0 {
			for i := 0; i < SHUFFLE_NUM; i++ {
				// 乱数でスワップする場所を決定
				sp1 := rand.Intn(len(rank_pool[rank]))
				sp2 := rand.Intn(len(rank_pool[rank]))
				// スワップ
				rank_pool[rank][sp1], rank_pool[rank][sp2] = rank_pool[rank][sp2], rank_pool[rank][sp1]
			}
		}
	}
	// 1つの配列にまとめる
	ret := []RegistrationInfo{}
	for rank := 0; rank < 5; rank++ {
		for i := 0; i < len(rank_pool[rank]); i++ {
			ret = append(ret, rank_pool[rank][i])
		}
	}
	return ret
}

func makeMatchPair(regi_list []RegistrationInfo) []Pair {
	ret := []Pair{}
	// １人余る場合は運営とマッチング
	if len(regi_list)%2 == 1 {
		// Todo: 末尾に運営を追加
	}
	for i := 0; i < len(regi_list); i += 2 {
		// 先頭から2人ペアを作成
		ret = append(ret, newPair(regi_list[i], regi_list[i+1]))
	}
	return ret
}

func makeMatchList(regi_list []Pair) []RegistrationInfo {
	ret := []RegistrationInfo{}
	for _, p := range regi_list {
		if p.First.user_id != -1 {
			ret = append(ret, p.First)
		}
		if p.Second.user_id != -1 {
			ret = append(ret, p.Second)
		}
	}
	return ret
}

func remakeList(latest_list []Pair, embed_list []RegistrationInfo) []Pair {
	ret := []RegistrationInfo{}
	latest := makeMatchList(latest_list)
	// rank 1から順に latest->embedの順でlistに入れる
	l_index := 0
	e_index := 0
	for r := 1; r <= 5; r++ {
		for l_index < len(latest) && latest[l_index].rank == r {
			ret = append(ret, latest[l_index])
			l_index++
		}
		for e_index < len(embed_list) && embed_list[e_index].rank == r {
			ret = append(ret, embed_list[e_index])
			e_index++
		}
	}
	return makeMatchPair(ret)
}

func session2_swapping(matching_list []Pair) []Pair {
	i := 0
	if len(matching_list)%2 == 1 {
		for ; i < len(matching_list)-3; i += 2 {
			matching_list[i].First, matching_list[i+1].First = matching_list[i+1].First, matching_list[i].First
		}
		// ラストは3組で入れ替え
		matching_list[i].First, matching_list[i+1].First, matching_list[i+2].First = matching_list[i+1].First, matching_list[i+2].First, matching_list[i].First
	} else {
		for ; i < len(matching_list); i += 2 {
			matching_list[i].First, matching_list[i+1].First = matching_list[i+1].First, matching_list[i].First
		}
	}
	return matching_list
}

func session3_swapping(matching_list []Pair) []Pair {
	i := 0
	if len(matching_list)%2 == 1 {
		for ; i < len(matching_list)-3; i += 2 {
			matching_list[i].First, matching_list[i+1].Second = matching_list[i+1].Second, matching_list[i].First
		}
		// ラストは3組で入れ替え
		matching_list[i].Second, matching_list[i+1].First, matching_list[i+1].Second, matching_list[i+2].First = matching_list[i+1].First, matching_list[i+2].First, matching_list[i].Second, matching_list[i+1].Second
	} else {
		for ; i < len(matching_list); i += 2 {
			matching_list[i].First, matching_list[i+1].Second = matching_list[i+1].Second, matching_list[i].First
		}
	}
	return matching_list
}
