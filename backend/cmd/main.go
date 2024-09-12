package main

import (
	"context"
	"log"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 環境変数を読み込む
	config.LoadEnv()

	// DSNを取得してデータベース接続を設定
	dsn := config.GetDSN()

	// Entクライアントを作成
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}
	defer client.Close()

	// データベーススキーマをマイグレーション
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// テストデータの挿入
	createTestData(client)
}

func createTestData(client *ent.Client) {
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
