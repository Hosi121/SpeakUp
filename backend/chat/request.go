package chat

// チャットの送信データ
type Request struct {
	Model     string            `json:"model"`      // モデルID
	Messages  []*RequestMessage `json:"messages"`   // メッセージのリスト
	MaxTokens int               `json:"max_tokens"` // 最大トークン数
}

func NewRequest(modelID string, messages []*RequestMessage, maxTokens int) *Request {
	return &Request{
		Model:     modelID,
		Messages:  messages,
		MaxTokens: maxTokens,
	}
}

// チャットの送信メッセージ
type RequestMessage struct {
	Role    string `json:"role"`    // メッセージの役割 (assistant, user, system)
	Content string `json:"content"` // メッセージの本文
}

func NewRequestMessage(role string, content string) *RequestMessage {
	return &RequestMessage{
		Role:    role,
		Content: content,
	}
}

