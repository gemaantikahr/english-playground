package wrapper

import (
	"english_playground/app/models/dto/request/send_notification"
	discord2 "english_playground/app/repositories/discord"
	request5 "english_playground/app/repositories/discord/models/request"
	fcm2 "english_playground/app/repositories/firebase/fcm"
	request3 "english_playground/app/repositories/firebase/fcm/models/request"
	otomat2 "english_playground/app/repositories/otomat"
	request6 "english_playground/app/repositories/otomat/models/request"
	sendtalk2 "english_playground/app/repositories/sendtalk"
	request2 "english_playground/app/repositories/sendtalk/models/request"
	"english_playground/app/repositories/telgeram"
	request4 "english_playground/app/repositories/telgeram/models/request"
	zoho2 "english_playground/app/repositories/zoho"
	"english_playground/app/repositories/zoho/models/request"
	"english_playground/app/response"
	"english_playground/app/services/discord"
	"english_playground/app/services/firebase/fcm"
	"english_playground/app/services/otomat"
	"english_playground/app/services/otomat/base"
	"english_playground/app/services/sendtalk"
	"english_playground/app/services/telegram"
	"english_playground/app/services/zoho"
	"fmt"
	"os"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

type handlerFunc func()

type WrapperChannelService struct {
	Request send_notification.SendNotificationDTORequest
}

func callFunctionOnAnotherThread(method handlerFunc) {
	defer time.Sleep(200 * time.Millisecond)
	method()
}

func (service WrapperChannelService) Do() response.ServiceResponse {

	specificIsNull := send_notification.SendNotificationDTORequest{}.Specific
	bulkIsNull := send_notification.SendNotificationDTORequest{}.Bulks

	// if specific request is not null then execute the process
	if !reflect.DeepEqual(service.Request.Specific, specificIsNull) {

		// if email exists then execute the process
		if service.Request.Specific.Email != nil {
			for _, email := range service.Request.Specific.Email {
				sendSpecificEmail(email.Recipient, email.Subject, email.Message)
			}
		}

		// if whatsapp exists then execute the process
		if service.Request.Specific.Wa != nil {
			for _, wa := range service.Request.Specific.Wa {
				sendSpecificWa(wa.Phone, wa.Message)
			}
		}

		// if fcm exists then execute the process
		if service.Request.Specific.Fcm != nil {
			for _, fcmToken := range service.Request.Specific.Fcm {
				sendSpecificFcm(fcmToken.Token, fcmToken.Routing, fcmToken.Id, fcmToken.Message, fcmToken.Title, fcmToken.Body, fcmToken.ImageUrl, fcmToken.Code)
			}
		}

		// if telegram exists then execute the process
		if service.Request.Specific.Telegram != nil {
			for _, tel := range service.Request.Specific.Telegram {
				sendSpecificTelegram(tel.Text, tel.ChatId)
			}
		}

		// if discord exists then execute the process
		if service.Request.Specific.Discord != nil {
			for _, dis := range service.Request.Specific.Discord {
				sendSpecificDiscord(dis.WebhookUrl, dis.Content)
			}
		}

	}

	// if bulk request is not null then execute the process
	if !reflect.DeepEqual(service.Request.Bulks, bulkIsNull) {

		emailIsNull := send_notification.SendNotificationDTORequest{}.Bulks.Email
		waIsNull := send_notification.SendNotificationDTORequest{}.Bulks.Wa
		fcmIsNull := send_notification.SendNotificationDTORequest{}.Bulks.Fcm
		discordIsNull := send_notification.SendNotificationDTORequest{}.Bulks.Discord

		// if email exists then execute the process
		if !reflect.DeepEqual(service.Request.Bulks.Email, emailIsNull) {
			sendBulkEmail(service.Request.Bulks.Email.Recipients, service.Request.Bulks.Email.Subject, service.Request.Bulks.Email.Message)
		}

		// if wa exists then execute the process
		if !reflect.DeepEqual(service.Request.Bulks.Wa, waIsNull) {
			sendBulkWa(service.Request.Bulks.Wa.Phones, service.Request.Bulks.Wa.Message)
		}

		// if fcm exists then execute the process
		if !reflect.DeepEqual(service.Request.Bulks.Fcm, fcmIsNull) {
			sendBulkFcmTokens(service.Request.Bulks.Fcm)
			sendBulkFcmTopic(service.Request.Bulks.Fcm)
		}

		// if discord exists then execute the process
		if !reflect.DeepEqual(service.Request.Bulks.Discord, discordIsNull) {
			sendBulkDiscord(service.Request.Bulks.Discord.WebhookUrls, service.Request.Bulks.Discord.Content)
		}

	}

	return response.ServiceResponse{}.Success("success", nil)

}

// function for send email to one recipient
func sendSpecificEmail(email string, subject string, message string) bool {

	// defined data for send email message
	var requestData request.SendMessageRequest
	requestData.Htmlbody = message
	requestData.Subject = subject
	requestData.From.Address = os.Getenv("ZOHO_SENDER_EMAIL")
	requestData.From.Name = os.Getenv("ZOHO_SENDER_NAME")
	// defined email address
	var emailAddress request.EmailAddress
	emailAddress.Address = email
	emailAddress.Name = "Pengguna Tauku"

	// defined to
	var to request.To
	to.EmailAddress = emailAddress
	requestData.To = append(requestData.To, to)

	callFunctionOnAnotherThread(func() {
		service := zoho.SendEmailService{
			Request: requestData,
			Repo:    zoho2.ZohoRepository{},
		}.Do()

		if !service.Status {
			log.Error("zoho.SendEmailService error ", service.Message)
		}

	})

	return true
}

// function for send wa to one recipient
func sendSpecificWa(phoneNumber string, message string) bool {

	callFunctionOnAnotherThread(func() {

		if os.Getenv("IS_OTOMAT") == "TRUE" {

			// send message using otomat
			var otomatRequestData request6.SendMessageRequest
			otomatRequestData.Phone = phoneNumber
			otomatRequestData.Text = message
			otomatService := otomat.SendMessageService{
				Repo: otomat2.OtomatRepository{Api: base.OtomatApi{
					OtomatConfig: base.OtomatConfig{},
					Type:         1,
				}},
				Request: otomatRequestData,
			}.Do()

			if !otomatService.Status {
				log.Error("otomat.SendWhatsappService error ", otomatService.Message)
			}

		}

		if os.Getenv("IS_SENDTALK") == "TRUE" {

			// send message using sendtalk
			var requestData request2.SendMessageRequest
			requestData.MessageType = "text"
			requestData.Phone = phoneNumber
			requestData.Body = message

			// call service
			service := sendtalk.SendWhatsappService{
				Request: requestData,
				Repo:    sendtalk2.SendTalkRepository{},
			}.Do()

			if !service.Status {
				log.Error("sendtalk.SendWhatsappService error ", service.Message)
			}

		}

	})

	return true

}

// func for send telegram to one chat
func sendSpecificTelegram(text string, chatId string) bool {

	// defined request data
	var sendMessageRequest request4.SendMessageRequest
	sendMessageRequest.Text = text
	sendMessageRequest.ChatId = chatId

	callFunctionOnAnotherThread(func() {

		// call service
		service := telegram.SendMessageService{
			Request: sendMessageRequest,
			Repo:    telgeram.TelegramRepository{},
		}.Do()

		if !service.Status {
			log.Error("sendtalk.sendSpecificTelegram error ", service.Message)
		}

	})

	return true

}

func sendSpecificDiscord(webhookUrl string, content string) bool {

	// defined request data
	var sendMessageRequest request5.SendMessageRequest
	sendMessageRequest.Content = content

	callFunctionOnAnotherThread(func() {

		// call service
		service := discord.SendMessageService{
			Repo:       discord2.DiscordRepository{},
			Request:    sendMessageRequest,
			WebhookUrl: webhookUrl,
		}.Do()

		if !service.Status {
			log.Error("discord.sendSpecificDiscord error ", service.Message)
		}
	})

	return true
}

// func for send fcm to one recipient
func sendSpecificFcm(token string, routing string, id string, message string, title string, body string, imageUrl string, code string) bool {

	var requestData request3.SendMessageRequest
	requestData.Message = message
	requestData.Id = id
	requestData.Routing = routing
	requestData.Title = title
	requestData.Body = body
	requestData.ImageUrl = imageUrl
	requestData.Code = code

	callFunctionOnAnotherThread(func() {

		service := fcm.SendMessageService{
			Request: requestData,
			Token:   token,
			Repo:    fcm2.FcmRepository{},
		}.Do()

		if !service.Status {
			log.Error("fcm.SendMessageService error ", service.Message)
		}
	})

	return true

}

// function for send email to multiple recipients
func sendBulkEmail(emails []string, subject string, message string) bool {

	var requestData request.SendMessageRequest
	requestData.Htmlbody = message
	requestData.Subject = subject
	requestData.From.Address = os.Getenv("ZOHO_SENDER_EMAIL")
	requestData.From.Name = os.Getenv("ZOHO_SENDER_NAME")

	for _, value := range emails {
		var emailAddress request.EmailAddress
		emailAddress.Address = value
		emailAddress.Name = "USER TAUKU"

		var to request.To
		to.EmailAddress.Address = value
		to.EmailAddress.Name = "USER TAUKU"

		requestData.To = append(requestData.To, to)

	}

	fmt.Println("emails", requestData)

	callFunctionOnAnotherThread(func() {
		// call service
		service := zoho.SendEmailService{
			Request: requestData,
			Repo:    zoho2.ZohoRepository{},
		}.Do()

		if !service.Status {
			log.Error("zoho.SendEmailService error", service.Message)
		}
	})

	return true
}

// function for send wa to multiple recipients
func sendBulkWa(numbers []string, message string) bool {

	for _, number := range numbers {

		callFunctionOnAnotherThread(func() {

			if os.Getenv("IS_OTOMAT") == "TRUE" {

				// send message using otomat
				var otomatRequestData request6.SendMessageRequest
				otomatRequestData.Phone = number
				otomatRequestData.Text = message
				otomatService := otomat.SendMessageService{
					Repo: otomat2.OtomatRepository{Api: base.OtomatApi{
						OtomatConfig: base.OtomatConfig{},
						Type:         1,
					}},
					Request: otomatRequestData,
				}.Do()

				if !otomatService.Status {
					log.Error("otomat.SendWhatsappService error ", otomatService.Message)
				}

			}

			if os.Getenv("IS_SENDTALK") == "TRUE" {

				// send message using sendtalk
				var requestData request2.SendMessageRequest
				requestData.MessageType = "text"
				requestData.Phone = number
				requestData.Body = message

				service := sendtalk.SendWhatsappService{
					Request: requestData,
					Repo:    sendtalk2.SendTalkRepository{},
				}.Do()

				if !service.Status {
					log.Error("sendtalk.SendWhatsappService error ", service.Message)
				}

			}

		})

	}

	return true

}

// function for send fcm to multiple recipients by token
func sendBulkFcmTokens(fcmBulk send_notification.FcmBulk) bool {

	for _, token := range fcmBulk.Token.Tokens {

		callFunctionOnAnotherThread(func() {

			var message request3.SendMessageRequest
			message.Message = fcmBulk.Token.Message
			message.Id = fcmBulk.Token.Id
			message.Routing = fcmBulk.Token.Routing
			message.Title = fcmBulk.Token.Title
			message.Body = fcmBulk.Token.Title
			message.ImageUrl = fcmBulk.Token.ImageUrl
			message.Code = fcmBulk.Token.Code

			// call service
			service := fcm.SendMessageService{
				Request: message,
				Token:   token,
				Repo:    fcm2.FcmRepository{},
			}.Do()

			if !service.Status {
				log.Error("fcm.SendMessageService error " + service.Message)
			}

		})

	}

	return true

}

// function for send fcm token to topic
func sendBulkFcmTopic(fcmBulk send_notification.FcmBulk) bool {

	callFunctionOnAnotherThread(func() {

		var message request3.SendMessageRequest
		message.Message = fcmBulk.Topic.Message
		message.Id = fcmBulk.Topic.Id
		message.Routing = fcmBulk.Topic.Routing
		message.Title = fcmBulk.Topic.Title
		message.Body = fcmBulk.Topic.Title
		message.ImageUrl = fcmBulk.Topic.ImageUrl
		message.Code = fcmBulk.Topic.Code

		// call service
		service := fcm.SendMessageService{
			Request: message,
			Topic:   fcmBulk.Topic.Name,
			Repo:    fcm2.FcmRepository{},
		}.Do()

		if !service.Status {
			log.Error("fcm.SendMessageService error " + service.Message)
		}

	})

	return true
}

func sendBulkDiscord(WebhookUrls []string, content string) bool {

	for _, webhookUrl := range WebhookUrls {

		var requestData request5.SendMessageRequest
		requestData.Content = content

		callFunctionOnAnotherThread(func() {

			// call service

			service := discord.SendMessageService{
				Repo:       discord2.DiscordRepository{},
				Request:    requestData,
				WebhookUrl: webhookUrl,
			}.Do()

			if !service.Status {
				log.Error("discord.sendBulkDiscord error ", service.Message)
			}
		})

	}

	return true
}
