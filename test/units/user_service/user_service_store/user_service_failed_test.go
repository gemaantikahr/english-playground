package user_service_store

import (
	"english_playground/app/models"
	"english_playground/app/response"
	"english_playground/app/services"
	"english_playground/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceStoreInternalServerError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	gomock.InOrder(
		mockRepo.EXPECT().Create(gomock.Eq("Dipa Ferdian")).Return(models.User{Id: 0, Name: ""}),
	)

	call := services.NewUser(mockRepo)
	actual := call.Store("Dipa Ferdian")

	expected := response.ServiceResponse{Status: false, Message: "Failed Insert Data", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}

func TestUserServiceStoreIsIntegerValueParameter(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	call := services.NewUser(mockRepo)
	actual := call.Store(1)

	expected := response.ServiceResponse{Status: false, Message: "Params name must be in string", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}

func TestUserServiceStoreIsEmptyValueParameter(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	call := services.NewUser(mockRepo)
	actual := call.Store(nil)

	expected := response.ServiceResponse{Status: false, Message: "Params name must be in string", Data: interface{}(nil)}

	assert.Equal(t, expected, actual)
}
