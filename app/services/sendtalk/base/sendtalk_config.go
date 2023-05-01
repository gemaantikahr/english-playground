package base

import "os"

type SendTalkConfig struct {
}

func (config SendTalkConfig) BaseUrl() string {
	return os.Getenv("SENDTALK_BASE_URL")
}

func (config SendTalkConfig) Token() string {
	return os.Getenv("SENDTALK_TOKEN")
}
