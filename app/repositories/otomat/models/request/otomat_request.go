package request

type SendMessageRequest struct {
	Phone string `json:"phone"`
	Text  string `json:"text"`
}
