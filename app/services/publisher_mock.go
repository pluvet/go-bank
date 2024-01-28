package services

import (
	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
)

var eventPublisher *eventpublisher.EventPublisher

func GetMockEventPublisher() *eventpublisher.EventPublisher {
	if eventPublisher == nil {
		var accountHandlers = map[string][]eventpublisher.Handler{
			new(events.EventUserCreated).GetName(): {},
		}
		eventPublisher = eventpublisher.NewEventPublisher(accountHandlers)
	}
	return eventPublisher
}
