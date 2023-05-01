package notification

import (
	"english_playground/app/models/dto/request/send_notification"
	"english_playground/app/models/dto/request/subscribe_topic"
	fcm2 "english_playground/app/repositories/firebase/fcm"
	"english_playground/app/repositories/firebase/fcm/models/request"
	"english_playground/app/response"
	"english_playground/app/services/firebase/fcm"
	"english_playground/app/services/wrapper"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
}

func (controller NotificationController) Send(ctx *gin.Context) {

	// binding request to send notification dto request
	var sendNotificationRequest send_notification.SendNotificationDTORequest
	if err := ctx.ShouldBind(&sendNotificationRequest); err != nil {
		ctx.JSON(400, response.Api().Error(err.Error(), nil))
		return
	}

	// call service
	wrapper.WrapperChannelService{
		Request: sendNotificationRequest,
	}.Do()

	ctx.JSON(200, response.Api().Success("success", sendNotificationRequest))

}

func (controller NotificationController) SubscribeTopic(ctx *gin.Context) {

	// binding request to subscribe topic dto request
	var subscribeTopicDTORequest subscribe_topic.SubscribeTopicDTORequest
	if err := ctx.ShouldBind(&subscribeTopicDTORequest); err != nil {
		ctx.JSON(400, response.Api().Error(err.Error(), nil))
		return
	}

	// defined subscribe topic request
	var subscribeTopicRequest request.SubScribeTopicRequest
	subscribeTopicRequest.Tokens = subscribeTopicDTORequest.Tokens
	subscribeTopicRequest.Name = subscribeTopicDTORequest.Name

	// call service
	service := fcm.SubscribeTopicService{
		Request: subscribeTopicRequest,
		Repo:    fcm2.FcmRepository{},
	}.Do()

	if !service.Status {
		ctx.JSON(400, response.Api().Error(service.Message, nil))
		return
	}

	ctx.JSON(200, response.Api().Success(service.Message, nil))

}
