package base

import "english_playground/app/base"

type ZohoApi struct {
	ZohoConfig
}

func (api ZohoApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(api.BaseUrl()+endpoint).AddHeader("Authorization", api.Token()).AddHeader("Content-Type", "application/json").AddHeader("Accept", "application/json")
}
