package handlers

import (
	"sync"

	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
)

type AccountHandler struct{}

func (a *AccountHandler) HandleEvent(event eventpublisher.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	eventUserCreated, ok := event.(*events.EventUserCreated)
	if ok {
		a.accountCreate(eventUserCreated.UserID)
	}
}

func (a *AccountHandler) accountCreate(userID int) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userID)
	config.DB.Create(&account)
}
