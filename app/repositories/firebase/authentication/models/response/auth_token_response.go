package response

type AuthTokenResponse struct {
	AuthTime int    `json:"auth_time"`
	Iss      string `json:"iss"`
	Aud      string `json:"aud"`
	Exp      int    `json:"exp"`
	Iat      int    `json:"iat"`
	Sub      string `json:"sub"`
	Uid      string `json:"uid"`
	Firebase struct {
		SignInProvider string `json:"sign_in_provider"`
		Tenant         string `json:"tenant"`
		Identities     struct {
			Email     []string `json:"email"`
			GoogleCom []string `json:"google.com"`
		} `json:"identities"`
	} `json:"firebase"`
}
