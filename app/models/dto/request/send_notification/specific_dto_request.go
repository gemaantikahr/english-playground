package send_notification

type EmailSpecific struct {
	Subject   string `json:"subject"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type WaSpecific struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type FcmSpecific struct {
	Message  string `json:"message"`
	Routing  string `json:"routing"`
	Id       string `json:"id"`
	Token    string `json:"token"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"image_url"`
	Code     string `json:"code"`
}

type TelegramSpecific struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type DiscordSpecific struct {
	WebhookUrl string `json:"webhook_url"`
	Content    string `json:"content"`
}
