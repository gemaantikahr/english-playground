package request

type SendMessageRequest struct {
	Phone       string `json:"phone"`
	MessageType string `json:"messageType"`
	Body        string `json:"body"`
}
