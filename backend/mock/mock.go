package mock

import (
	"context"
	"log"

	"github.com/Hosi121/SpeakUp/ent"
)

func CreateTestData(client *ent.Client) {
	ctx := context.Background()

	// ユーザーの作成
	user, err := client.USERS.
		Create().
		SetUsername("testuser").
		SetEmail("testuser@example.com").
		SetHashedPassword("hashedpassword").
		SetAvatarURL("https://example.com/avatar.png").
		SetRole("USER").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	log.Printf("created user: %v", user)
}
