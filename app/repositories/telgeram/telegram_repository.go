package telgeram

import (
	"encoding/json"
	"english_playground/app/repositories/telgeram/models/request"
	"english_playground/app/repositories/telgeram/models/response"
	"english_playground/app/services/telegram/base"
	"io"
)

type TelegramRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type TelegramRepository struct {
}

func (repo TelegramRepository) SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error) {

	// defined endpoint
	endpoint := "/bot" + base.TelegramConfig{}.Token() + "/sendMessage"
	call, err := base.TelegramApi{}.Post(endpoint).Bodys(request).Call()
	if err != nil {
		return response.SendMessageResponse{}, err
	}

	// read response body
	result, _ := io.ReadAll(call.Body)
	var sendMessageResponse response.SendMessageResponse

	// map response body
	if errMap := json.Unmarshal(result, &sendMessageResponse); errMap != nil {
		return response.SendMessageResponse{}, errMap
	}

	return sendMessageResponse, nil
}
