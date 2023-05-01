package response

type SendMessageResponse struct {
	Data []struct {
		Code           string        `json:"code"`
		AdditionalInfo []interface{} `json:"additional_info"`
		Message        string        `json:"message"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Object    string `json:"object"`
}
