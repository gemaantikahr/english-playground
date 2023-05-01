package response

type SendMessageResponse struct {
	Status int    `json:"status"`
	ReqID  string `json:"reqID"`
	Error  struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"error"`
	Data struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
		Id      string `json:"id"`
	} `json:"data"`
}
