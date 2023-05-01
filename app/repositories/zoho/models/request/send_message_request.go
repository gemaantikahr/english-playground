package request

type SendMessageRequest struct {
	From     From   `json:"from"`
	To       []To   `json:"to"`
	Subject  string `json:"subject"`
	Htmlbody string `json:"htmlbody"`
}

type From struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type To struct {
	EmailAddress EmailAddress `json:"email_address"`
}

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
