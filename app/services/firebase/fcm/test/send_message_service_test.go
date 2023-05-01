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

func TestSendMessageService_Do(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined data
	var message request.SendMessageRequest
	message.Message = "you are pukiman"
	message.Id = "1"
	message.Routing = ":*"
	message.ImageUrl = "https://storage.googleapis.com/mskrjowrjowijo24-tauku-mkdjwkej/images/logo-name.png"
	message.Title = "the title"
	message.Body = "the body"

	token := "fDBdKmTYT3W3JtHjp45stt:APA91bGtsi2vetb47Pp_syLmNZkw92hkuc-TP_EBJuPU7QDqQ_b6rB4sAvMdvFUX7JqfB1EZ5qW6T9KB9XJA_clvUDKqcBM_q7SCdcuTUerDFuaiUjH3L93pllWcKUq-cuSbA6pxwwTk"
	topic := ""
	// call service
	service := fcm.SendMessageService{
		Request: message,
		Token:   token,
		Topic:   topic,
		Repo:    fcm2.FcmRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("success", service)

}
