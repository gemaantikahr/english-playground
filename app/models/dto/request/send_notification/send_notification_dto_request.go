package send_notification

type SendNotificationDTORequest struct {
	Specific struct {
		Email    []EmailSpecific    `json:"email"`
		Wa       []WaSpecific       `json:"wa"`
		Fcm      []FcmSpecific      `json:"fcm"`
		Telegram []TelegramSpecific `json:"telegram"`
		Discord  []DiscordSpecific  `json:"discord"`
	} `json:"specific"`
	Bulks struct {
		Email    EmailBulk    `json:"email"`
		Wa       WaBulk       `json:"wa"`
		Fcm      FcmBulk      `json:"fcm"`
		Telegram TelegramBulk `json:"telegram"`
		Discord  DiscordBulk  `json:"discord"`
	} `json:"bulks"`
}
