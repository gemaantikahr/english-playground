package base

import (
	"english_playground/app/base"
)

type OtomatApi struct {
	OtomatConfig
	Type int
}

func (api OtomatApi) Get(endpoint string) base.NetClient {

	if api.Type == 1 {

		data := map[string]string{
			"api_id":  api.ApiID1(),
			"api_key": api.ApiKey1(),
		}

		baseUrl := api.BaseUrl1()

		return base.HttpService().Get().Url(baseUrl + endpoint).AddUrlParam(data)

	} else {

		data := map[string]string{
			"api_id":  api.ApiID2(),
			"api_key": api.ApiKey2(),
		}

		baseUrl := api.BaseUrl2()

		return base.HttpService().Get().Url(baseUrl + endpoint).AddUrlParam(data)
	}

}
