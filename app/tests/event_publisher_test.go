package tests

import (
	"testing"

	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
	"github.com/pluvet/go-bank/app/publisher"
	"github.com/stretchr/testify/assert"
)

func TestGetEventPublisher(t *testing.T) {
	eventPublisher := publisher.GetEventPublisher()
	var accountHandlers = map[string][]eventpublisher.Handler{
		new(events.EventUserCreated).GetName(): {new(handlers.AccountHandler)},
	}

	assert.Equal(t, eventpublisher.NewEventPublisher(accountHandlers), eventPublisher)
}

func TestEventPublisherSentEvent(t *testing.T) {
	eventPublisher := publisher.GetEventPublisher()

	var eventUserCreated = events.NewEventUserCreated(1)
	go eventPublisher.NewEvent(eventUserCreated)
}
