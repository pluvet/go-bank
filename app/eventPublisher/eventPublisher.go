package eventPublisher

import (
	"sync"
)

type Event interface {
	GetName() string
}

type Handler interface {
	ReactEvent(Event, *sync.WaitGroup)
}

type EventPublisher struct {
	EventName string
}

func NewEventPublisher() *EventPublisher {
	var eventPublisher = new(EventPublisher)
	return eventPublisher
}

func (e *EventPublisher) NewEvent(handlers map[string][]Handler, event Event) {
	var wg sync.WaitGroup
	e.EventName = event.GetName()
	eventHandlers := handlers[e.EventName]
	for i := range eventHandlers {
		wg.Add(1)
		var handler = eventHandlers[i]
		go handler.ReactEvent(event, &wg)
	}
	wg.Wait()
}
