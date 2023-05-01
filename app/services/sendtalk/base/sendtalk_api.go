package base

import "english_playground/app/base"

type SendTalkApi struct {
	SendTalkConfig
}

func (api SendTalkApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(api.BaseUrl()+endpoint).AddHeader("API-Key", api.Token())
}
