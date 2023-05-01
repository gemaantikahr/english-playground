package test

import (
	sendtalk2 "english_playground/app/repositories/sendtalk"
	"english_playground/app/repositories/sendtalk/models/request"
	"english_playground/app/services/sendtalk"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSendWhatsappService_Do(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	for i := 1; i <= 5; i++ {

		i := i
		go func() {
			// defined data
			var requestData request.SendMessageRequest
			requestData.MessageType = "text"
			requestData.Phone = "6289680988"
			requestData.Body = "awkwk ini bot " + strconv.Itoa(i)

			// call service
			service := sendtalk.SendWhatsappService{
				Request: requestData,
				Repo:    sendtalk2.SendTalkRepository{},
			}.Do()

			if !service.Status {
				println("something went wrong ! ", service.Message)
				return
			}

			log.WithFields(log.Fields{
				"response": service.Data,
			}).Info("SEND WHATSAPP SERVICE RESULT")

		}()
	}

	time.Sleep(5 * time.Second)

}
