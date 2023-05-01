package base

import "english_playground/app/base"

type DiscordApi struct {
}

func (api DiscordApi) Post(webhookUrl string) base.NetClient {
	return base.HttpService().Post().Url(webhookUrl).AddHeader("Content-Type", "application/json")
}
