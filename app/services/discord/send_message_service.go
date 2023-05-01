package discord

import (
	"english_playground/app/repositories/discord"
	"english_playground/app/repositories/discord/models/request"
	"english_playground/app/response"
)

type SendMessageService struct {
	Repo       discord.DiscordRepoInterface
	Request    request.SendMessageRequest
	WebhookUrl string
}

func (service SendMessageService) Do() response.ServiceResponse {

	// call repo
	repo, err := service.Repo.SendMessage(service.Request, service.WebhookUrl)
	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	if repo.Error.Message != "" {
		return response.Service().Error(repo.Error.Message, nil)
	}

	return response.Service().Success("success", repo)

}
