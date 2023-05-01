package test

import (
	"english_playground/app/repositories/discord"
	"english_playground/app/repositories/discord/models/request"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestDiscordRepository_SendMessage(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined message request
	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.Content = "lovely message from golang"

	// defined webhook url
	webhookUrl := "https://discord.com/api/webhooks/1076385115246186496/gnguBOeMdEp2mpEolK6AWlDLa_LHotUh2hn0VDkC_fG3MZNNF58V92KHVj8Gf-xGIL46"

	// call repository

	repo, errRepo := discord.DiscordRepository{}.SendMessage(sendMessageRequest, webhookUrl)
	if errRepo != nil {
		println("something went wrong ! ", errRepo.Error())
		return
	}

	fmt.Println("success", repo)

}
