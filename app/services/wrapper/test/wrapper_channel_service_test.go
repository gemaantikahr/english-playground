package test

import (
	"english_playground/app/models/dto/request/send_notification"
	"english_playground/app/services/wrapper"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestWrapperChannelService_SpecificEmail(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined email
	var email1 send_notification.EmailSpecific
	email1.Subject = "testing email 1 subject"
	email1.Recipient = "gemaantikahr@gmail.com"
	email1.Message = "testing email 1 message"

	var email2 send_notification.EmailSpecific
	email2.Subject = "testing email 2 subject"
	email2.Recipient = "gema1600018095@webmail.uad.ac.id"
	email2.Message = "testing email 2 message"

	// defined specific email
	var specific send_notification.SendNotificationDTORequest
	specific.Specific.Email = append(specific.Specific.Email, email2, email1)

	log.WithFields(log.Fields{
		"response": specific,
	}).Info("SPECIFIC DATA")

	service := wrapper.WrapperChannelService{Request: specific}.Do()

	println("result ", service.Message)
}

func TestWrapperChannelService_SpecificWa(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined wa
	var wa1 send_notification.WaSpecific
	wa1.Phone = "6289680988232"
	wa1.Message = "hey you pukiman, keep sprite"

	var wa2 send_notification.WaSpecific
	wa2.Phone = "6281313989234"
	wa2.Message = "hey you pukiman, keep sprite drunk"

	// defined specific wa
	var specificWa send_notification.SendNotificationDTORequest
	specificWa.Specific.Wa = append(specificWa.Specific.Wa, wa1, wa2)

	// call the service
	service := wrapper.WrapperChannelService{Request: specificWa}.Do()
	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	println("success")

}

func TestWrapperChannelService_SpecificFcm(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined fcm
	var fcm1 send_notification.FcmSpecific
	fcm1.Token = "c2kR8AHSt09Kp6SpqHO4jG:APA91bHoviYq9dEPhWH5e_QknTmg2DimW2tDHRakpU9ezROJ2YJTye5Vfw6K6YiuKHiI5jBhp9hQeWV2BT3Uk7UIJERmFeNq5XynwUhJ9nuNeCDPd2inOj_hHrDTB-H5Gd4yE0mE1l0o"
	fcm1.Message = "testing message for you"

	// defined specific fcm
	var requestData send_notification.SendNotificationDTORequest
	requestData.Specific.Fcm = append(requestData.Specific.Fcm, fcm1)

	// call the service
	service := wrapper.WrapperChannelService{Request: requestData}.Do()
	if !service.Status {
		println("something went wrong ! ")
		return
	}

	println("success ", service.Message)
}

func TestWrapperChannelService_BulkEmail(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined send notification
	var sendNotification send_notification.SendNotificationDTORequest

	var emails = []string{"gemaantikahr@gmail.com", "gema1600018095@webmail.uad.ac.id"}

	// defined email
	var emailBulk send_notification.EmailBulk
	emailBulk.Recipients = emails
	emailBulk.Message = "this is bulk message"
	emailBulk.Subject = "subject email"

	sendNotification.Bulks.Email = emailBulk

	service := wrapper.WrapperChannelService{
		Request: sendNotification,
	}.Do()

	log.WithFields(log.Fields{
		"request data": sendNotification,
	}).Info("SPECIFIC DATA")

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	println("success")

}

func TestWrapperChannelService_BulkWa(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined phone
	phones := []string{"6289680988232"}
	// defined bulk wa
	var waBulk send_notification.WaBulk
	waBulk.Phones = phones
	waBulk.Message = "keep sprite pukiman hero"

	// defined send message request
	var messageRequest send_notification.SendNotificationDTORequest
	messageRequest.Bulks.Wa = waBulk

	// call service
	service := wrapper.WrapperChannelService{Request: messageRequest}.Do()
	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	println("successes")

}

func TestWrapperChannelService_BulkFcm(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	var tokens []string
	for i := 1; i < 2; i++ {
		tokens = append(tokens, "fpRWOsluSHWbliB9OTzkTB:APA91bFBmXLlNhnbDz-jAmwc0H1JNbt0erXRR9xLLFgELSOJsu9IBZ4o_Y9g1KGlrmzbE2XVsqIKKQluFerIQm00-MEzfmSDZHBFRnj5t5fMCgFon0-QkWwy0sUA1XullsYDp9abCnFW")
	}

	// defined fcm token
	var fcm send_notification.FcmBulk
	fcm.Token.Tokens = tokens
	fcm.Token.Id = "1"
	fcm.Token.Message = "detail jos"
	fcm.Token.Routing = "/order_detail_fnb"
	fcm.Token.Title = "title from test class"
	fcm.Token.Body = "body from test class"
	fcm.Token.ImageUrl = "https://storage.googleapis.com/mskrjowrjowijo24-tauku-mkdjwkej/images/logo-name.png"

	// defined request data
	var sendNotificationRequest send_notification.SendNotificationDTORequest
	sendNotificationRequest.Bulks.Fcm = fcm

	// call service
	service := wrapper.WrapperChannelService{Request: sendNotificationRequest}.Do()
	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	println(service.Message)

}

func TestWrapperChannelService_SpecificTelegram(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined first chat id
	var firstTelegramSpecific send_notification.TelegramSpecific
	firstTelegramSpecific.ChatId = "658555117"
	firstTelegramSpecific.Text = "god only one !"

	var wrapperRequest send_notification.SendNotificationDTORequest
	wrapperRequest.Specific.Telegram = append(wrapperRequest.Specific.Telegram, firstTelegramSpecific)

	// call service
	sendMessageService := wrapper.WrapperChannelService{wrapperRequest}.Do()
	if !sendMessageService.Status {
		println("something went wrong !")
		return
	}

	fmt.Println("success ", sendMessageService)
}
