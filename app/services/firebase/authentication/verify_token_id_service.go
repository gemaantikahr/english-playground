package authentication

import (
	"english_playground/app/repositories/firebase/authentication"
	"english_playground/app/response"
)

type VerifyTokenIDService struct {
	TokenID string
	Repo    authentication.AuthenticationRepoInterface
}

func (service VerifyTokenIDService) Do() response.ServiceResponse {

	// call repository
	repo, errRepo := service.Repo.VerifyIDToken(service.TokenID)

	if errRepo != nil {
		return response.ServiceResponse{}.Error(errRepo.Error(), nil)
	}

	return response.ServiceResponse{}.Success("success", repo)

}
