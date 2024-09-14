package main

import (
	"context"
	"log"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/routes"

	middleware "github.com/Hosi121/SpeakUp/middlewares"
	supabaseAPI "github.com/Hosi121/SpeakUp/supaseAPI"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 環境変数を読み込む
	config.LoadEnv()

	// DSNを取得してデータベース接続を設定
	dsn := config.GetDSN()

	// ginの初期化
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SupabaseAuthRoutes(r)

	// supabase APIの準備
	supabaseAPI.InitSupabase()

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
	// mock.CreateTestData(client)

	// 8080ポートでサーバーを起動
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
