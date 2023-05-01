package test

import (
	"english_playground/app/repositories/zoho"
	"english_playground/app/repositories/zoho/models/request"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestZohoRepository_SendMessage(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("cannot load env : ", err)
	}

	// defined request data
	var requestData request.SendMessageRequest
	requestData.Htmlbody = "testing message body"
	requestData.Subject = "testing subject"
	requestData.From.Address = os.Getenv("ZOHO_SENDER_EMAIL")
	requestData.From.Name = os.Getenv("ZOHO_SENDER_NAME")
	// defined email address
	var emailAddress request.EmailAddress
	emailAddress.Address = "gemaantikahr@gmail.com"
	emailAddress.Name = "gema antika hariadi"

	// defined to
	var to request.To
	to.EmailAddress = emailAddress
	requestData.To = append(requestData.To, to)

	// call repository
	repo, errRepo := zoho.ZohoRepository{}.SendMessage(requestData)
	if errRepo != nil {
		println("something went wrong ", errRepo.Error())
		return
	}

	fmt.Println("repo response ", repo)

}
