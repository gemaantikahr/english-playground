package fcm

import (
	"english_playground/app/repositories/firebase/fcm"
	"english_playground/app/repositories/firebase/fcm/models/request"
	"english_playground/app/response"
)

type SubscribeTopicService struct {
	Request request.SubScribeTopicRequest
	Repo    fcm.FcmRepoInterface
}

func (service SubscribeTopicService) Do() response.ServiceResponse {

	// call repo
	repo, err := service.Repo.SubscribeTopic(service.Request)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	return response.Service().Success(repo, nil)

}
