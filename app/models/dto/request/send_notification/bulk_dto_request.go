package send_notification

type EmailBulk struct {
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Message    string   `json:"message"`
}

type WaBulk struct {
	Phones  []string `json:"phones"`
	Message string   `json:"message"`
}

type FcmBulk struct {
	Topic struct {
		Name     string `json:"name"`
		Message  string `json:"message"`
		Routing  string `json:"routing"`
		Id       string `json:"id"`
		Title    string `json:"title"`
		Body     string `json:"body"`
		ImageUrl string `json:"image_url"`
		Code     string `json:"code"`
	} `json:"topic"`
	Token struct {
		Tokens   []string `json:"tokens"`
		Message  string   `json:"message"`
		Routing  string   `json:"routing"`
		Id       string   `json:"id"`
		Title    string   `json:"title"`
		Body     string   `json:"body"`
		ImageUrl string   `json:"image_url"`
		Code     string   `json:"code"`
	} `json:"token"`
}

type TelegramBulk struct {
	ChannelId string `json:"channel_id"`
	Message   string `json:"message"`
}

type DiscordBulk struct {
	WebhookUrls []string `json:"webhook_urls"`
	Content     string   `json:"content"`
}
