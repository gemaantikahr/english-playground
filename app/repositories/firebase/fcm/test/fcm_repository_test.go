package test

import (
	"context"
	"english_playground/app/base"
	"english_playground/app/repositories/firebase/fcm"
	"english_playground/app/repositories/firebase/fcm/models/request"
	"fmt"
	"os"
	"testing"

	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestFcmRepository_SendMessage(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined data
	var messageRequest request.SendMessageRequest
	messageRequest.Message = "yay, you got an order !"
	messageRequest.Routing = "/order_detail_fnb"
	messageRequest.Id = "267"

	// token
	token := "fpRWOsluSHWbliB9OTzkTB:APA91bFBmXLlNhnbDz-jAmwc0H1JNbt0erXRR9xLLFgELSOJsu9IBZ4o_Y9g1KGlrmzbE2XVsqIKKQluFerIQm00-MEzfmSDZHBFRnj5t5fMCgFon0-QkWwy0sUA1XullsYDp9abCnFW"
	topic := "first-topic"

	response, errRepo := fcm.FcmRepository{}.SendMessage(messageRequest, token, topic)
	if errRepo != nil {
		println("error", errRepo.Error())
		return
	}

	fmt.Println("succeess", response)

}

func TestSubscribeTopic_Fcm(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	client, errM := base.InitFcm()
	if errM != nil {
		println("something went wrong ! ", errM.Error())
		return
	}

	registrationTokens := []string{
		"fpRWOsluSHWbliB9OTzkTB:APA91bFBmXLlNhnbDz-jAmwc0H1JNbt0erXRR9xLLFgELSOJsu9IBZ4o_Y9g1KGlrmzbE2XVsqIKKQluFerIQm00-MEzfmSDZHBFRnj5t5fMCgFon0-QkWwy0sUA1XullsYDp9abCnFW",
	}

	ctx := context.Background()
	topic := "first-topic"
	// Subscribe the devices corresponding to the registration tokens to the
	// topic.
	response, err := client.SubscribeToTopic(ctx, registrationTokens, topic)
	if err != nil {
		log.Fatalln(err)
	}
	// See the TopicManagementResponse reference documentation
	// for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
}

func TestSendMessageToTopic(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	client, errM := base.InitFcm()
	if errM != nil {
		println("something went wrong ! ", errM.Error())
		return
	}

	notification := &messaging.Notification{
		Title:    "sipaling title",
		Body:     "sipaling body",
		ImageURL: "https://storage.googleapis.com/mskrjowrjowijo24-tauku-mkdjwkej/image/1672643660.jpg",
	}

	message := &messaging.Message{
		Data: map[string]string{
			"message": "mantap",
			"routing": "no routing",
			"id":      "yuhu",
		},
		Topic:        "first-topic",
		Notification: notification,
	}

	response, errSend := client.Send(context.Background(), message)
	if errSend != nil {
		println("something went wrong ! ", errSend.Error())
		return
	}

	log.WithFields(log.Fields{
		"response": response,
	}).Info("FCM CLASS SEND MESSAGE TO TOPIC")

}

func TestFcmRepository_SubscribeTopic(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined topic
	topic := "kamu_ganteng"
	tokens := []string{
		"fpRWOsluSHWbliB9OTzkTB:APA91bFBmXLlNhnbDz-jAmwc0H1JNbt0erXRR9xLLFgELSOJsu9IBZ4o_Y9g1KGlrmzbE2XVsqIKKQluFerIQm00-MEzfmSDZHBFRnj5t5fMCgFon0-QkWwy0sUA1XullsYDp9abCnFW",
		"XXXX",
	}

	// defined request data
	var requestData request.SubScribeTopicRequest
	requestData.Tokens = tokens
	requestData.Name = topic

	// call repo
	repo, errRepo := fcm.FcmRepository{}.SubscribeTopic(requestData)
	if errRepo != nil {
		println("something went wrong ! ", errRepo.Error())
		return
	}

	println("result ", repo)

}
