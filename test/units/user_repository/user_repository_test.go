package user_repository

import (
	"english_playground/app/models"
	"english_playground/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	mockRepo.EXPECT().Find(gomock.Eq(1)).Return(models.User{Id: 1, Name: "Dipa Ferdian"})

	call := mockRepo.Find(1)

	expect := models.User{
		Id:   1,
		Name: "Dipa Ferdian",
	}

	assert.Equal(t, expect, call)
}
