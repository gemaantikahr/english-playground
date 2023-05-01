package request

type SendMessageRequest struct {
	Text   string `json:"text"`
	ChatId string `json:"chat_id"`
}
