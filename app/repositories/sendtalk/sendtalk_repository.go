package sendtalk

import (
	"encoding/json"
	"english_playground/app/repositories/sendtalk/models/request"
	"english_playground/app/repositories/sendtalk/models/response"
	"english_playground/app/services/sendtalk/base"
	"io"

	log "github.com/sirupsen/logrus"
)

type SendTalkRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type SendTalkRepository struct {
}

func (repo SendTalkRepository) SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error) {

	// defined send message response
	var sendMessageResponse response.SendMessageResponse

	// call send wa api
	call, err := base.SendTalkApi{}.Post("/v1/message/send_whatsapp").Bodys(request).Call()
	if err != nil {
		return sendMessageResponse, err
	}

	// read api response body
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return sendMessageResponse, errRa
	}

	// map response body to send message response
	if errMap := json.Unmarshal(result, &sendMessageResponse); errMap != nil {
		return sendMessageResponse, errMap
	}

	log.WithFields(log.Fields{
		"response": sendMessageResponse,
	}).Info("SEND TALK RESPONSE")

	return sendMessageResponse, nil

}
