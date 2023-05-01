package base

import "os"

type TelegramConfig struct {
}

func (config TelegramConfig) BaseUrl() string {
	return os.Getenv("TELEGRAM_BASE_URL")
}

func (config TelegramConfig) Token() string {
	return os.Getenv("TELEGRAM_TOKEN")
}
