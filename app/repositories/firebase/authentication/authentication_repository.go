package authentication

import (
	"context"
	"english_playground/app/base"
	"english_playground/app/repositories/firebase/authentication/models/response"
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type AuthenticationRepoInterface interface {
	VerifyIDToken(tokenID string) (*auth.Token, error)
	GetUserByEmail(email string) (*auth.UserRecord, error)
}

type AuthenticationRepository struct {
}

func (repo AuthenticationRepository) VerifyIDToken(tokenID string) (*auth.Token, error) {

	// variable token
	var tokenResponse *auth.Token

	// initialize authentication
	client, errInit := base.InitAuth()
	if errInit != nil {
		return tokenResponse, errInit
	}

	// verifying token
	tokenVer, errVer := client.VerifyIDToken(context.Background(), tokenID)
	if errVer != nil {
		return tokenVer, errVer
	}

	var authTokenResponse response.AuthTokenResponse

	if errMap := mapstructure.Decode(tokenVer, &authTokenResponse); errMap != nil {
		return tokenVer, errMap
	}

	fmt.Println("oke doke ", authTokenResponse.Firebase.Identities.Email)

	for _, value := range authTokenResponse.Firebase.Identities.Email {
		println("value ", value)
	}

	log.WithFields(log.Fields{
		"response": tokenVer,
	}).Info("VERIFY TOKEN ID RESULT")

	return tokenVer, nil
}

func (repo AuthenticationRepository) GetUserByEmail(email string) (*auth.UserRecord, error) {

	// initial user record
	var userRecord *auth.UserRecord

	// initialize client
	client, errInit := base.InitAuth()
	if errInit != nil {
		return userRecord, errInit
	}

	user, errUser := client.GetUserByEmail(context.Background(), email)
	if errUser != nil {
		return userRecord, errUser
	}

	log.WithFields(log.Fields{
		"response": user,
	}).Info("GET USER BY EMAIL RESULT")

	return user, nil
}
