package test

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"testing"
)

func TestOauth(t *testing.T) {

	opt := option.WithCredentialsFile("../gemakey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		println("something went wrong ! ", err.Error())
		return
	}

	client, err := app.Auth(context.Background())

	fmt.Println("value", client)
}

func TestKedua(t *testing.T) {

	opt := option.WithCredentialsFile("../tauku_dev_personal.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		println("something wrong", err.Error())
		return
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		println("something wrong", err.Error())
		return
	}

	// This registration token comes from the client FCM SDKs.

	// See documentation on defining a message payload.
	//registrationToken := "ffBIhQC5yGhCJ-0LU0K5Ay:APA91bEI1LhYwIXMzUhVGgvLvMH4JBAFFHBE7TW3otNY-XYHQl4Ri1J4qztlkk20L6q6QZb1i65wKFau4RGZi3IZrYVrywX8sbEQvuOm_LhjFtRR8njz35BRzTFJ0V-SAE-QR0sz5xwh"
	registrationToken := "d9DP3uxrGdK8eBi9gdhwNh:APA91bHtMZQOg9lWd08cBpthTwGTNtkzfNdQvs-3cGSGoKv_JLG-RzUPcHAn3RE4wD9JFdshkH9sEqbXletbx-CZ_MMDjgzk3-ZLz5tv6gTK2IKndxPQWtpvln25cCyq6JdaxGOzoahv"
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		println("something wrong", err.Error())
		return
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
