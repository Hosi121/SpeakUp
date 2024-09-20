package mock

import (
	"context"
	"log"
	"strconv"

	"github.com/Hosi121/SpeakUp/ent"
)

func CreateTestData(client *ent.Client) {
	ctx := context.Background()

	// ユーザーの作成
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
		}

		log.Printf("created user: %v", user)
	}
	admin, err := client.USERS.
		Create().
		SetUsername("admin").
		SetEmail("admin@example.com").
		SetAvatarURL("https://example.com/avatar.png").
		SetRole("ADMIN").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating admin: %v", err)
	}
	log.Printf("created user: %v", admin)

}
