package services

import (
	"fmt"

	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/repositories"
)

type UserService struct {
	repo           repositories.IUserRepository
	eventPublisher eventpublisher.EventPublisher
}

func NewUserService(repo repositories.IUserRepository, eventPublisher eventpublisher.EventPublisher) *UserService {
	userService := new(UserService)
	userService.repo = repo
	userService.eventPublisher = eventPublisher
	return userService
}

func (u *UserService) CreateUser(name string, email string, password string) (*int, error) {

	userID, err := u.repo.CreateUser(name, email, password)

	if err != nil {
		return nil, err
	}

	var eventUserCreated = events.NewEventUserCreated(*userID)
	eventWasPublished := u.eventPublisher.NewEvent(eventUserCreated)

	if !eventWasPublished {
		fmt.Printf("eventUserCreated was not published")
	}

	return userID, nil
}
