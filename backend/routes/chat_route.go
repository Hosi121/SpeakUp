package routes

import (
	"log"
	"time"

	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/Hosi121/SpeakUp/utils"
	"github.com/gin-gonic/gin"
)

func ChatRoute(router *gin.Engine, apiKey string) {
	// JSONファイルからプロンプトをロード
	prompt, err := utils.LoadPromptConfig("prompt.json")
	if err != nil {
		log.Fatal("Error loading prompt config:", err)
	}

	// ChatControllerの初期化時にプロンプトを設定
	chatController := controllers.NewChatController(
		"gpt-3.5-turbo", // 使用するモデル
		apiKey,          // 環境変数から取得したAPIキー
		100,             // max tokens
		20*time.Second,  // timeout
		prompt,          // JSONから読み込んだプロンプトを渡す
	)

	chatGroup := router.Group("/chat")
	{
		chatGroup.POST("/ask", chatController.AskQuestion)
	}
}
