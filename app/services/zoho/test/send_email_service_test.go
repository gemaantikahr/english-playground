package test

import (
	zoho2 "english_playground/app/repositories/zoho"
	"english_playground/app/repositories/zoho/models/request"
	"english_playground/app/services/zoho"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSendEmailService_Do(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	var requestData request.SendMessageRequest
	requestData.Htmlbody = "yes yo do"
	requestData.Subject = "NEW ORDER"
	requestData.From.Address = os.Getenv("ZOHO_SENDER_EMAIL")
	requestData.From.Name = os.Getenv("ZOHO_SENDER_NAME")
	// defined email address
	var emailAddress request.EmailAddress
	emailAddress.Address = "gemaantikahr@gmail.com"
	emailAddress.Name = "gema antika hariadi"

	// defined email address
	var emailAddress1 request.EmailAddress
	emailAddress1.Address = "gema1600018095@webmail.uad.ac.id"
	emailAddress1.Name = "gema antika hariadi"

	// defined to
	var to request.To
	to.EmailAddress = emailAddress

	// defined to 1
	var to1 request.To
	to1.EmailAddress = emailAddress1

	requestData.To = append(requestData.To, to, to1)

	// call service
	service := zoho.SendEmailService{
		Request: requestData,
		Repo:    zoho2.ZohoRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ! ")
		return
	}

	println("success")
}
