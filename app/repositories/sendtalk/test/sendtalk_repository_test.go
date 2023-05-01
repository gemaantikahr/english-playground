package test

import (
	"english_playground/app/repositories/sendtalk"
	"english_playground/app/repositories/sendtalk/models/request"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSendTalkRepository_SendMessage(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined data
	var requestData request.SendMessageRequest
	requestData.MessageType = "text"
	requestData.Phone = "6289680988232"
	requestData.Body = "testing gema"

	// call repo
	_, err = sendtalk.SendTalkRepository{}.SendMessage(requestData)

	fmt.Println("error", err)
}
