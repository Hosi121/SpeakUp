package controllers

import (
	"net/http"
	"time"

	"github.com/Hosi121/SpeakUp/chat"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatCompletions *chat.ChatCompletions
	prompt          string // 事前に設定されたプロンプト
}

// NewChatController は新しい ChatController を作成します
func NewChatController(model string, secret string, maxTokens int, timeout time.Duration, prompt string) *ChatController {
	chatCompletions := chat.NewChatCompletions(model, secret, maxTokens, timeout)
	return &ChatController{
		chatCompletions: chatCompletions,
		prompt:          prompt, // 事前設定されたプロンプトを保存
	}
}

// AskQuestion は1つの質問に対する応答を返します
func (cc *ChatController) AskQuestion(c *gin.Context) {
	// フロントエンドから送られてきた質問文を取得
	var request struct {
		Content string `json:"content"` // メインの質問文
	}

	// リクエストボディをJSONとしてバインド
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 事前設定されたプロンプトと質問文を組み合わせる
	finalMessage := cc.prompt + "\n" + request.Content

	// Chat APIに送信するメッセージを作成
	requestMessage := chat.NewRequestMessage("user", finalMessage)
	response, err := cc.chatCompletions.SendMessage([]*chat.RequestMessage{requestMessage})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response"})
		return
	}

	// 応答をJSONで返す
	c.JSON(http.StatusOK, response)
}
