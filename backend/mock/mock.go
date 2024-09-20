package mock

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Hosi121/SpeakUp/ent"
)

func CreateTestData(client *ent.Client) {
	ctx := context.Background()

	// イベントの作成
	theme, err := client.AI_THEMES.
		Create().
		SetThemeText("hogehoge").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating AI theme: %v", err)
	} else {
		log.Printf("created AI theme: %v", theme)
	}

	event, err := client.EVENTS.
		Create().
		SetEventStart(time.Date(2024, 9, 21, 15, 0, 0, 0, time.Local)).
		SetEventEnd(time.Date(2024, 9, 21, 16, 0, 0, 0, time.Local)).
		SetThemeID(theme.ID).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating event: %v", err)
	} else {
		log.Printf("created event: %v", event)
	}

	event2, err := client.EVENTS.
		Create().
		SetEventStart(time.Date(2024, 9, 31, 19, 0, 0, 0, time.Local)).
		SetEventEnd(time.Date(2024, 9, 31, 20, 0, 0, 0, time.Local)).
		SetThemeID(theme.ID).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating event: %v", err)
	} else {
		log.Printf("created event: %v", event2)
	}

	// ユーザーの作成＆イベント参加登録
	for i := 1; i <= 30; i++ {
		tail := strconv.Itoa(i)
		user, err := client.USERS.
			Create().
			SetUsername("testuser" + tail).
			SetEmail("testuser" + tail + "@example.com").
			SetAvatarURL("https://example.com/avatar.png").
			SetRole("USER").
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		} else {
			log.Printf("created user: %v", user)
		}

		// イベント参加登録
		event_record, err := client.EVENT_RECORDS.
			Create().
			SetUserID(user.ID).
			SetEventID(event.ID).
			SetParticipatesBit(rand.Intn(7) + 1).
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating record: %v", err)
		} else {
			log.Printf("created record: %v", event_record)
		}
	}
	admin, err := client.USERS.
		Create().
		SetUsername("admin").
		SetEmail("admin@example.com").
		SetAvatarURL("https://example.com/avatar.png").
		SetRole("ADMIN").
		SetRank(rand.Intn(5) + 1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating admin: %v", err)
	} else {
		log.Printf("created user: %v", admin)
	}

}
