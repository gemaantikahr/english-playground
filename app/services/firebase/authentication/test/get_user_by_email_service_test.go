package test

import (
	authentication2 "english_playground/app/repositories/firebase/authentication"
	"english_playground/app/services/firebase/authentication"
	"fmt"
	"os"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func TestGetUserByEmailService(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// call service
	service := authentication.GetUserByEmailService{
		Email: "yeremia997@gmail.com",
		Repo:  authentication2.AuthenticationRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ", service.Message)
		return
	}

	// mapped response
	var userRecord *auth.UserRecord

	if errMap := mapstructure.Decode(service.Data, &userRecord); errMap != nil {
		println("something went wrong ! ", errMap.Error())
		return
	}

	fmt.Println("success with email", userRecord.DisplayName)

}
