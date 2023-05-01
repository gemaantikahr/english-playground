package test

import (
	discord2 "english_playground/app/repositories/discord"
	request2 "english_playground/app/repositories/discord/models/request"
	"english_playground/app/services/discord"
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
	request.Content = "message from service class test"

	// defined webhook url
	webhookUrl := "https://discord.com/api/webhooks/1076385115246186496/gnguBOeMdEp2mpEolK6AWlDLa_LHotUh2hn0VDkC_fG3MZNNF58V92KHVj8Gf-xGIL46"

	// call service

	service := discord.SendMessageService{
		Repo:       discord2.DiscordRepository{},
		Request:    request,
		WebhookUrl: webhookUrl,
	}.Do()

	if !service.Status {
		println("something went wrong !")
		return
	}

	fmt.Println("success ", service.Data)
}
