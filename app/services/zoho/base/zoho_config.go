package base

import "os"

type ZohoConfig struct {
}

func (config ZohoConfig) BaseUrl() string {
	return os.Getenv("ZOHO_BASE_URL")
}

func (config ZohoConfig) Token() string {
	return os.Getenv("ZOHO_TOKEN")
}

func (config ZohoConfig) SenderEmail() string {
	return os.Getenv("ZOHO_SENDER_EMAIL")
}

func (config ZohoConfig) SenderName() string {
	return os.Getenv("ZOHO_SENDER_NAME")
}

func (config ZohoConfig) OtpBody() string {
	return os.Getenv("ZOHO_OTP_BODY")
}
