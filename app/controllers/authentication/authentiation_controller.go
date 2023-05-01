package authentication

import (
	"english_playground/app/models/dto/request/authentication"
	authentication3 "english_playground/app/repositories/firebase/authentication"
	"english_playground/app/response"
	authentication2 "english_playground/app/services/firebase/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
}

func (controller AuthenticationController) VerifyToken(ctx *gin.Context) {

	// binding data to AuthenticationDTORequest
	var authenticationDTORequest authentication.AuthenticationDTORequest
	if err := ctx.ShouldBind(&authenticationDTORequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(err.Error(), nil))
		return
	}

	// call service
	verifyTokenService := authentication2.ExecuteVerifyTokenService{
		TokenId: authenticationDTORequest.TokenId,
		Repo:    authentication3.AuthenticationRepository{},
	}.Do()

	if !verifyTokenService.Status {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(verifyTokenService.Message, nil))
		return
	}

	ctx.JSON(http.StatusOK, response.Api().Success("success", verifyTokenService.Data))
	return
}
