package test

import (
	fcm2 "english_playground/app/repositories/firebase/fcm"
	"english_playground/app/repositories/firebase/fcm/models/request"
	"english_playground/app/services/firebase/fcm"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSubscribeTopicService_Do(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined data
	topic := "kamu_cakep"
	tokens := []string{
		"fpRWOsluSHWbliB9OTzkTB:APA91bFBmXLlNhnbDz-jAmwc0H1JNbt0erXRR9xLLFgELSOJsu9IBZ4o_Y9g1KGlrmzbE2XVsqIKKQluFerIQm00-MEzfmSDZHBFRnj5t5fMCgFon0-QkWwy0sUA1XullsYDp9abCnFW",
		"XXXX",
	}

	// defined request
	var requestData request.SubScribeTopicRequest
	requestData.Name = topic
	requestData.Tokens = tokens

	// call service
	service := fcm.SubscribeTopicService{
		Request: requestData,
		Repo:    fcm2.FcmRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	println("success ", service.Message)
}
