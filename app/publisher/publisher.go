package publisher

import (
	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
)

var EventPublisher *eventpublisher.EventPublisher

func Init() {
	var eventUserCreated = new(events.EventUserCreated)
	accountHandler := new(handlers.AccountHandler)
	var accountHandlers = map[string][]eventpublisher.Handler{
		eventUserCreated.GetName(): {accountHandler},
	}
	EventPublisher = eventpublisher.NewEventPublisher(accountHandlers)
}
