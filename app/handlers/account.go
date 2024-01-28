package handlers

import (
	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/repositories"
	"github.com/pluvet/go-bank/app/services"
)

type AccountHandler struct {
}

func (a *AccountHandler) HandleEvent(event eventpublisher.Event) error {
	eventUserCreated, ok := event.(*events.EventUserCreated)

	if !ok {
		err := new(ErrorEventIsNotSupported)
		return err
	}

	accountService := services.NewAccountService(new(repositories.AccountRepository), nil)
	_, err := accountService.CreateAccount(eventUserCreated.UserID)
	if err != nil {
		return err
	}
	return nil
}
