package test

import (
	otomat2 "english_playground/app/repositories/otomat"
	"english_playground/app/repositories/otomat/models/request"
	"english_playground/app/services/otomat"
	"english_playground/app/services/otomat/base"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSendMessageService_Do(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined message
	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.Phone = "6289680988232"
	sendMessageRequest.Text = "here message from testing class"

	// call service
	service := otomat.SendMessageService{
		Repo: otomat2.OtomatRepository{Api: base.OtomatApi{
			OtomatConfig: base.OtomatConfig{},
			Type:         1,
		}},
		Request: sendMessageRequest,
	}.Do()

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("success ", service.Data)

}
