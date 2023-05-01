package base

import "english_playground/app/base"

type TelegramApi struct {
	TelegramConfig
}

func (api TelegramApi) Post(endpoint string) base.NetClient {
	return base.NetClient{}.Post().Url(api.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}
