package handlers

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/eventPublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
)

type AccountHandler struct {
	UserID int
}

func NewAccountHandler() *AccountHandler {
	var a = new(AccountHandler)
	return a
}

func (a *AccountHandler) ReactEvent(event eventPublisher.Event) {
	switch event := event.(type) {
	case *events.EventUserCreated:
		a.UserID = event.UserID
	}
	a.AccountCreate()
}

func (a *AccountHandler) AccountCreate() {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(a.UserID)
	config.DB.Create(&account)
}
