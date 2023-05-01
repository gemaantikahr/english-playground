package telegram

import (
	"english_playground/app/repositories/telgeram"
	"english_playground/app/repositories/telgeram/models/request"
	"english_playground/app/response"
)

type SendMessageService struct {
	Request request.SendMessageRequest
	Repo    telgeram.TelegramRepoInterface
}

func (service SendMessageService) Do() response.ServiceResponse {

	// call repo
	repo, errRepo := service.Repo.SendMessage(service.Request)
	if errRepo != nil {
		return response.Service().Error(errRepo.Error(), nil)
	}

	return response.Service().Success("success", repo)

}
