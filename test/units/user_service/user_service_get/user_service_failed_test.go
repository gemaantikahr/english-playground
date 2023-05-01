package user_service_get

import (
	"english_playground/app/models"
	"english_playground/app/response"
	"english_playground/app/services"
	"english_playground/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	gomock.InOrder(
		mockRepo.EXPECT().Find(gomock.Eq(1)).Return(models.User{Id: 0, Name: ""}),
	)

	call := services.NewUser(mockRepo)
	actual := call.Get(1)
	expected := response.ServiceResponse{Status: false, Message: "user not found", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}

func TestUserServiceIsStringValueParameter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	call := services.NewUser(mockRepo)
	actual := call.Get("a")

	expected := response.ServiceResponse{Status: false, Message: "Params must be in integer", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}

func TestUserServiceIsEmptyValueParameter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	call := services.NewUser(mockRepo)
	actual := call.Get(nil)

	expected := response.ServiceResponse{Status: false, Message: "Params must be in integer", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}
