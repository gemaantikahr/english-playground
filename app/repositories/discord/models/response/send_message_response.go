package response

type SendMessageResponse struct {
	Error Error `json:"error"`
}
type Error struct {
	Message   string   `json:"message"`
	Code      int      `json:"code"`
	WebhookId []string `json:"webhook_id"`
}
