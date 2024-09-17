package chat

// チャットの受信データ
type Response struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   *Usage    `json:"usage"`
	Choices []*Choice `json:"choices"`
}

// APIの使用量
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      *ResponseMessage `json:"message"`
	FinishReason string           `json:"finish_reason"`
	Index        int              `json:"index"`
}

// チャットの受信メッセージ
type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

