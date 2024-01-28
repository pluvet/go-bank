package eventpublisher

import (
	"testing"
	"time"

	"github.com/pluvet/go-bank/app/events"
	"github.com/stretchr/testify/assert"
)

type HandlerTest struct {
}

func (a *HandlerTest) HandleEvent(event Event) error {
	time.Sleep(35 * time.Second)
	return nil
}

func GetMockEventPublisher() *EventPublisher {
	var accountHandlers = map[string][]Handler{
		new(events.EventUserCreated).GetName(): {},
	}
	eventPublisher := NewEventPublisher(accountHandlers)

	return eventPublisher
}

func GetMockEventPublisherWithOneHandler() *EventPublisher {
	var accountHandlers = map[string][]Handler{
		new(events.EventAccountBalanceIncreased).GetName(): {new(HandlerTest)},
	}
	eventPublisherWithOneHandler := NewEventPublisher(accountHandlers)

	return eventPublisherWithOneHandler
}

func TestEventPublisherSentEventSuccess(t *testing.T) {
	eventPublisher := GetMockEventPublisher()

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(12, 0)

	eventWasPublished := eventPublisher.NewEvent(eventAccountBalanceIncreased)
	assert.Equal(t, true, eventWasPublished)
}

func TestEventPublisherProcessEventSuccess(t *testing.T) {
	eventPublisherWithOneHandler := GetMockEventPublisherWithOneHandler()

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(12, 0)
	eventWasPublished := make(chan bool)
	go eventPublisherWithOneHandler.processEvent(eventAccountBalanceIncreased, eventWasPublished)
	result := <-eventWasPublished
	assert.Equal(t, true, result)
}
