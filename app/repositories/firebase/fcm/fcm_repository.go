package fcm

import (
	"context"
	"english_playground/app/base"
	"english_playground/app/repositories/firebase/fcm/models/request"

	"firebase.google.com/go/messaging"
	log "github.com/sirupsen/logrus"
)

type FcmRepoInterface interface {
	SendMessage(request request.SendMessageRequest, token string, topic string) (string, error)
	SubscribeTopic(request request.SubScribeTopicRequest) (string, error)
}

type FcmRepository struct {
}

func (repo FcmRepository) SubscribeTopic(request request.SubScribeTopicRequest) (string, error) {

	// init fcm app
	client, errM := base.InitFcm()
	if errM != nil {
		return "", errM
	}

	response, errS := client.SubscribeToTopic(context.Background(), request.Tokens, request.Name)
	if errS != nil {
		return "", errS
	}

	log.WithFields(log.Fields{
		"response": response,
	}).Info("FCM REPOSITORY SUBSCRIBE TOPIC")

	return "success", nil
}

func (repo FcmRepository) SendMessage(request request.SendMessageRequest, token string, topic string) (string, error) {

	// init fcm app
	client, errM := base.InitFcm()
	if errM != nil {
		return "", errM
	}

	var notificationTitle string
	var notificationBody string
	var notificationImageUrl string

	if request.Body != "" {
		notificationBody = request.Body
	} else {
		notificationBody = request.Message
	}

	if request.Title != "" {
		notificationTitle = request.Title
	} else {
		notificationTitle = "New notification"
	}

	if request.ImageUrl != "" {
		notificationImageUrl = request.ImageUrl
	}

	notification := &messaging.Notification{
		Title:    notificationTitle,
		Body:     notificationBody,
		ImageURL: notificationImageUrl,
	}

	// defined message data
	message := &messaging.Message{}

	// defined message for token notification
	if token != "" {

		message = &messaging.Message{
			Data: map[string]string{
				"message": request.Message,
				"routing": request.Routing,
				"id":      request.Id,
				"code":    request.Code,
			},
			Token:        token,
			Notification: notification,
		}

	}

	// defined message for topic notification
	if topic != "" {

		message = &messaging.Message{
			Data: map[string]string{
				"message": request.Message,
				"routing": request.Routing,
				"id":      request.Id,
				"code":    request.Code,
			},
			Topic:        topic,
			Notification: notification,
		}

	}

	// send message
	response, errSend := client.Send(context.Background(), message)
	if errSend != nil {
		return "", errSend
	}

	log.WithFields(log.Fields{
		"response": repo,
	}).Info("FCM REPOSITORY SEND MESSAGE")

	return response, nil

}
