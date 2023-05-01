package zoho

import (
	"encoding/json"
	"english_playground/app/repositories/zoho/models/request"
	"english_playground/app/repositories/zoho/models/response"
	"english_playground/app/services/zoho/base"
	"io"
)

type ZohoRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type ZohoRepository struct {
}

func (repo ZohoRepository) SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error) {

	// defined send message response
	var sendMessageResponse response.SendMessageResponse
	// call send message api
	call, errApi := base.ZohoApi{}.Post("/v1.1/email").Bodys(request).Call()
	if errApi != nil {
		return sendMessageResponse, errApi
	}

	// read response body
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return sendMessageResponse, errApi
	}

	// mapping response body
	if errMap := json.Unmarshal(result, &sendMessageResponse); errMap != nil {
		return sendMessageResponse, errMap
	}

	return sendMessageResponse, nil

}
