package utils

import (
	"context"
	"log/slog"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/achievements"
)

func Awarding() {
	// DBの用意
	dsn := config.GetDSN()
	db_client, err := ent.Open("mysql", dsn)
	if err != nil {
		slog.Error("Failed to open connection to database: %v", err)
	}
	defer db_client.Close()

	ctx := context.Background()

	// PROGRESSテーブルを全て取得
	progresses, err := db_client.PROGRESS.Query().All(ctx)
	if err != nil {
		slog.Error("Failed querying progresses: %v", err)
		return
	}
	// 1人ずつトロフィーの条件を満たしているか確認
	for _, progress := range progresses {
		// 各トロフィーの条件判定
		// Todo: 各トロフィーのtitle決め
		switch progress.LoginDays { // ログイン日数に関するトロフィー
		case 1:
			checkAchievement(progress.UserID, "First Login", db_client, ctx)
		case 7:
			checkAchievement(progress.UserID, "1 week Login", db_client, ctx)
		case 30:
			checkAchievement(progress.UserID, "1 month Login", db_client, ctx)
		}
		switch progress.ConsecutiveLoginDays { // 連続ログイン日数に関するトロフィー
		case 7:
			checkAchievement(progress.UserID, "1 week Sequence Login", db_client, ctx)
		case 30:
			checkAchievement(progress.UserID, "1 month Sequence Login", db_client, ctx)
		}
	}
}

func checkAchievement(user_id int, title string, db_client *ent.Client, ctx context.Context) {
	// 該当ユーザが既にトロフィーを獲得していないか確認
	_, err := db_client.ACHIEVEMENTS.
		Query().
		Where(
			achievements.UserIDEQ(user_id),
			achievements.TitleEQ(title),
		).
		Only(ctx)
	if err != nil { // Not Found or Error
		if !ent.IsNotFound(err) {
			slog.Error("Failed querying progresses: %v", err)
			return
		}
	}
	// トロフィーを付与
	_, err2 := db_client.ACHIEVEMENTS.
		Create().
		SetUserID(user_id).
		SetTitle(title).
		Save(ctx)
	if err2 != nil {
		slog.Error("Failed creating achievement: %v", err2)
	}
}
