package routes

import (
	"log"
	"time"

	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/Hosi121/SpeakUp/utils"
	"github.com/gin-gonic/gin"
)

func ChatRoute(router *gin.Engine, apiKey string) {
	// theme_prompt.json をロード
	themePrompt, err := utils.LoadPromptConfig("theme_prompt.json")
	if err != nil {
		log.Fatal("Error loading theme prompt config:", err)
	}

	// support_prompt.json をロード
	supportPrompt, err := utils.LoadPromptConfig("support_prompt.json")
	if err != nil {
		log.Fatal("Error loading support prompt config:", err)
	}

	// ChatControllerの初期化時にプロンプトを設定
	themeChatController := controllers.NewChatController(
		"gpt-3.5-turbo", // 使用するモデル
		apiKey,          // 環境変数から取得したAPIキー
		100,             // max tokens
		20*time.Second,  // timeout
		themePrompt,     // JSONから読み込んだプロンプトを渡す (theme)
	)

	supportChatController := controllers.NewChatController(
		"gpt-3.5-turbo", // 使用するモデル
		apiKey,          // 環境変数から取得したAPIキー
		1000,            // max tokens
		20*time.Second,  // timeout
		supportPrompt,   // JSONから読み込んだプロンプトを渡す (support)
	)

	chatGroup := router.Group("/chat")
	{
		// /theme エンドポイントには theme_prompt.json を使用
		chatGroup.POST("/theme", themeChatController.AskQuestion)

		// /ask エンドポイントには support_prompt.json を使用
		chatGroup.POST("/ask", supportChatController.AskQuestion)
	}
}
