package zoho

import (
	"english_playground/app/repositories/zoho"
	"english_playground/app/repositories/zoho/models/request"
	"english_playground/app/response"
)

type SendEmailService struct {
	Request request.SendMessageRequest
	Repo    zoho.ZohoRepoInterface
}

func (service SendEmailService) Do() response.ServiceResponse {

	// call repo
	repo, errRepo := service.Repo.SendMessage(service.Request)
	if errRepo != nil {
		return response.Service().Error(errRepo.Error(), nil)
	}

	return response.Service().Success("success", repo)

}
