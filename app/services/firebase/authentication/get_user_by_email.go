package authentication

import (
	"english_playground/app/repositories/firebase/authentication"
	"english_playground/app/response"
)

type GetUserByEmailService struct {
	Email string
	Repo  authentication.AuthenticationRepoInterface
}

func (service GetUserByEmailService) Do() response.ServiceResponse {

	// call repo
	userRepo, errRepo := service.Repo.GetUserByEmail(service.Email)
	if errRepo != nil {
		return response.ServiceResponse{}.Error(errRepo.Error(), nil)
	}

	return response.ServiceResponse{}.Success("success", userRepo)

}
