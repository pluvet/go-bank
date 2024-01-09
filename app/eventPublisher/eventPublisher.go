package eventPublisher

import (
	"sync"
)

type Event interface {
	GetName() string
}

type Handler interface {
	HandleEvent(Event, *sync.WaitGroup)
}

type EventPublisher struct {
	handlers map[string][]Handler
}

func NewEventPublisher(handlers map[string][]Handler) *EventPublisher {
	var eventPublisher = new(EventPublisher)
	eventPublisher.handlers = handlers
	return eventPublisher
}

func (e *EventPublisher) NewEvent(event Event) {
	var wg sync.WaitGroup
	eventHandlers := e.handlers[event.GetName()]
	for i := range eventHandlers {
		wg.Add(1)
		var handler = eventHandlers[i]
		go handler.HandleEvent(event, &wg)
	}
	wg.Wait()
}
