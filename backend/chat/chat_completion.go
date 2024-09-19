package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"io"
	"net/http"
)

type ChatCompletions struct {
	// HTTPリクエストのタイムアウト
	timeout time.Duration

	// 応答の最大トークン
	maxTokens int

	// チャットに使用するモデルのID
	model string

	// APIキー
	secret string
}

func NewChatCompletions(model string, secret string, maxTokens int, timeout time.Duration) *ChatCompletions {
	return &ChatCompletions{
		maxTokens: maxTokens,
		model:     model,
		secret:    secret,
		timeout:   timeout,
	}
}

// APIにメッセージを送信する
func (c ChatCompletions) SendMessage(messages []*RequestMessage) (*Response, error) {
	data, err := json.Marshal(NewRequest(c.model, messages, c.maxTokens))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.secret)

	client := &http.Client{
		// リソース節約のためにタイムアウトを設定する
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("bad status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// 一回きりの質問をする
func (c ChatCompletions) AskOneQuestion(content string) (*Response, error) {
	messages := []*RequestMessage{
		NewRequestMessage("user", content),
	}
	return c.SendMessage(messages)
}



