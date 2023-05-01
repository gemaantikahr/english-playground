package test

import (
	"english_playground/app/repositories/telgeram"
	request2 "english_playground/app/repositories/telgeram/models/request"
	"english_playground/app/services/telegram"
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

	// defined request data
	var request request2.SendMessageRequest
	request.Text = "testing"
	request.ChatId = "658555117"
	// call service
	service := telegram.SendMessageService{
		Request: request,
		Repo:    telgeram.TelegramRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("success ", service.Message)
}
