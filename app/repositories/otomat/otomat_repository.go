package otomat

import (
	"encoding/json"
	"english_playground/app/helpers"
	"english_playground/app/repositories/otomat/models/request"
	"english_playground/app/repositories/otomat/models/response"
	"english_playground/app/services/otomat/base"
	"io"
)

type OtomatRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type OtomatRepository struct {
	Api base.OtomatApi
}

func (repo OtomatRepository) SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error) {

	// convert send message request to params as map data type
	param := helpers.ConvertStructToMapString{Input: request}.Do()

	// call api
	call, err := repo.Api.Get("").AddUrlParam(param).Call()
	if err != nil {
		return response.SendMessageResponse{}, err
	}

	// read response body
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return response.SendMessageResponse{}, errRa
	}

	// mapping result to send message response
	var sendMessageResponse response.SendMessageResponse
	if errMap := json.Unmarshal(result, &sendMessageResponse); errMap != nil {
		return sendMessageResponse, errMap
	}

	return sendMessageResponse, nil

}
