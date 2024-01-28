package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pluvet/go-bank/app/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateService(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := repositories.NewMockIUserRepository(ctrl)

	id := 1
	m.
		EXPECT().
		CreateUser("test name", "test@mail.com", "123456").
		Return(&id, nil).
		Times(1)
	eventpublisher := GetMockEventPublisher()
	userService := NewUserService(m, *eventpublisher)
	userID, err := userService.CreateUser("test name", "test@mail.com", "123456")
	assert.Equal(t, &id, userID)
	assert.NoError(t, err)
}
