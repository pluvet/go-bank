package publisher

import (
	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
)

var eventPublisher *eventpublisher.EventPublisher

func GetEventPublisher() *eventpublisher.EventPublisher {
	if eventPublisher == nil {
		var accountHandlers = map[string][]eventpublisher.Handler{
			new(events.EventUserCreated).GetName(): {new(handlers.AccountHandler)},
		}
		eventPublisher = eventpublisher.NewEventPublisher(accountHandlers)
	}
	return eventPublisher
}
