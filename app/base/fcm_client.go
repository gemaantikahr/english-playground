package base

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"os"
)

func InitFcm() (*messaging.Client, error) {

	opt := option.WithCredentialsFile(os.Getenv("FCM_SERVICE_ACCOUNT_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, errM := app.Messaging(ctx)
	if errM != nil {
		return nil, errM
	}

	return client, nil

}

func InitAuth() (*auth.Client, error) {

	opt := option.WithCredentialsFile(os.Getenv("FCM_SERVICE_ACCOUNT_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, errM := app.Auth(ctx)
	if errM != nil {
		return nil, errM
	}

	return client, nil
}
