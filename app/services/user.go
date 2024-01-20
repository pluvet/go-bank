package services

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
)

func CreateUser(user *models.User) error {

	result := config.DB.Create(&user)

	if result.Error != nil {
		err := new(ErrorCreatingRecordInDB)
		err.Model = "user"
		return err
	}

	var eventUserCreated = events.NewEventUserCreated(user.ID)
	eventPublisher := publisher.GetEventPublisher()
	go eventPublisher.NewEvent(eventUserCreated)

	return nil
}
