package fcm

import (
	"english_playground/app/repositories/firebase/fcm"
	"english_playground/app/repositories/firebase/fcm/models/request"
	"english_playground/app/response"
)

type SendMessageService struct {
	Request request.SendMessageRequest
	Token   string
	Topic   string
	Repo    fcm.FcmRepoInterface
}

func (service SendMessageService) Do() response.ServiceResponse {

	// call repo
	repo, errRepo := service.Repo.SendMessage(service.Request, service.Token, service.Topic)

	if errRepo != nil {
		return response.Service().Error(errRepo.Error(), nil)
	}

	return response.Service().Success("success", repo)

}
