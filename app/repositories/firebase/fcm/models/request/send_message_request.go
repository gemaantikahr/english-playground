package request

type SendMessageRequest struct {
	Message  string `json:"message"`
	Routing  string `json:"routing"`
	Id       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"image_url"`
	Code     string `json:"code"`
}
