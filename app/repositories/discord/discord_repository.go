package discord

import (
	"encoding/json"
	"english_playground/app/repositories/discord/models/request"
	"english_playground/app/repositories/discord/models/response"
	"english_playground/app/services/discord/base"
	"io"

	log "github.com/sirupsen/logrus"
)

type DiscordRepoInterface interface {
	SendMessage(request request.SendMessageRequest, webhookUrl string) (response.SendMessageResponse, error)
}

type DiscordRepository struct {
}

func (repo DiscordRepository) SendMessage(request request.SendMessageRequest, webhookUrl string) (response.SendMessageResponse, error) {

	// call discord webhook api
	call, err := base.DiscordApi{}.Post(webhookUrl).Bodys(request).Call()
	if err != nil {
		return response.SendMessageResponse{}, err
	}

	// read api response body
	result, _ := io.ReadAll(call.Body)

	// map response body to send message response
	var sendMessageResponse response.SendMessageResponse
	if errMap := json.Unmarshal(result, &sendMessageResponse); errMap != nil {
		return sendMessageResponse, errMap
	}

	log.WithFields(log.Fields{
		"response": sendMessageResponse,
	}).Info("DISCORD RESPONSE")

	return sendMessageResponse, nil

}
