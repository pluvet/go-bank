package publisher

import (
	"github.com/pluvet/go-bank/app/eventPublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
)

var EVENTPUBLISHER *eventPublisher.EventPublisher

func Init() {
	var eventUserCreated = new(events.EventUserCreated)
	accountHandler := new(handlers.AccountHandler)
	var accountHandlers = map[string][]eventPublisher.Handler{
		eventUserCreated.GetName(): {accountHandler},
	}
	eventPublisher := eventPublisher.NewEventPublisher(accountHandlers)
	EVENTPUBLISHER = eventPublisher
}
