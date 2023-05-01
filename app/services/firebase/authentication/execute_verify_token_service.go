package authentication

import (
	"english_playground/app/repositories/firebase/authentication"
	response2 "english_playground/app/repositories/firebase/authentication/models/response"
	"english_playground/app/response"
	response3 "english_playground/app/services/firebase/authentication/models/response"

	"github.com/mitchellh/mapstructure"
)

type ExecuteVerifyTokenService struct {
	TokenId string
	Repo    authentication.AuthenticationRepoInterface
}

func (service ExecuteVerifyTokenService) Do() response.ServiceResponse {

	// call verify token repo
	tokenRepo, errRepo := service.Repo.VerifyIDToken(service.TokenId)
	if errRepo != nil {
		return response.ServiceResponse{}.Error(errRepo.Error(), nil)
	}

	var authTokenResponse response2.AuthTokenResponse
	if errMap := mapstructure.Decode(tokenRepo, &authTokenResponse); errMap != nil {
		return response.ServiceResponse{}.Error(errMap.Error(), nil)
	}

	// get email from verify token response
	var email string
	for _, value := range authTokenResponse.Firebase.Identities.Email {
		email = value
	}

	// call get user record
	userRecord, errRepoUser := service.Repo.GetUserByEmail(email)
	if errRepoUser != nil {
		return response.ServiceResponse{}.Error(errRepoUser.Error(), nil)
	}

	// initialize service response
	var executeVerifyTokenServiceResponse response3.ExecuteVerifyTokenServiceResponse
	executeVerifyTokenServiceResponse.PhotoUrl = userRecord.PhotoURL
	executeVerifyTokenServiceResponse.PhoneNumber = userRecord.PhoneNumber
	executeVerifyTokenServiceResponse.Uid = userRecord.UID
	executeVerifyTokenServiceResponse.ProviderId = userRecord.ProviderID
	executeVerifyTokenServiceResponse.UserDisplay = userRecord.DisplayName
	executeVerifyTokenServiceResponse.Email = userRecord.Email

	return response.ServiceResponse{}.Success("success", executeVerifyTokenServiceResponse)

}
