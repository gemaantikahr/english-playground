package sendtalk

import (
	"english_playground/app/repositories/sendtalk"
	"english_playground/app/repositories/sendtalk/models/request"
	"english_playground/app/response"
)

type SendWhatsappService struct {
	Request request.SendMessageRequest
	Repo    sendtalk.SendTalkRepoInterface
}

func (service SendWhatsappService) Do() response.ServiceResponse {

	// call the repo
	repo, err := service.Repo.SendMessage(service.Request)

	if err != nil {
		return response.ServiceResponse{}.Error(err.Error(), nil)
	}

	return response.ServiceResponse{}.Success("success", repo)

}
