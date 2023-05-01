package test

import (
	"english_playground/app/repositories/otomat"
	"english_playground/app/repositories/otomat/models/request"
	"english_playground/app/services/otomat/base"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestOtomatRepository_SendMessage(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined message request
	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.Phone = "6289680988232"
	sendMessageRequest.Text = "testing banget"

	// defined config
	repo, errRepo := otomat.OtomatRepository{Api: base.OtomatApi{
		OtomatConfig: base.OtomatConfig{},
		Type:         1,
	}}.SendMessage(sendMessageRequest)

	if errRepo != nil {
		println("something went wrong ! ", errRepo.Error())
		return
	}

	fmt.Println("success ", repo)

}
