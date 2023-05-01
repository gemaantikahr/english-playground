package response

type ExecuteVerifyTokenServiceResponse struct {
	UserDisplay string `json:"user_display"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PhotoUrl    string `json:"photo_url"`
	ProviderId  string `json:"provider_id"`
	Uid         string `json:"uid"`
}
