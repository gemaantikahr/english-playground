package otomat

import (
	"english_playground/app/repositories/otomat"
	"english_playground/app/repositories/otomat/models/request"
	"english_playground/app/response"
)

type SendMessageService struct {
	Repo    otomat.OtomatRepoInterface
	Request request.SendMessageRequest
}

func (service SendMessageService) Do() response.ServiceResponse {

	// call repo
	repo, errRepo := service.Repo.SendMessage(service.Request)
	if errRepo != nil {
		return response.Service().Error(errRepo.Error(), nil)
	}

	return response.Service().Success("success", repo)

}
