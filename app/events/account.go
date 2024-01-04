package events

import (
	"github.com/pluvet/go-bank/app/handlers"
)

type Event interface {
	CreateAccountEvent()
}

type AccountEvent struct {
	Event
	handler handlers.Handler
	userID  uint
}

func NewEvent(handler handlers.Handler, userID uint) Event {
	var a = new(AccountEvent)
	a.handler = handler
	a.userID = userID
	return a
}

func (a *AccountEvent) CreateAccountEvent() {
	go a.handler.AccountCreate(a.userID)
}
