package eventPublisher

type Event interface {
	GetName() string
}

type Handler interface {
	ReactEvent(Event)
}

type EventPublisher struct {
	EventName string
}

func NewEventPublisher() *EventPublisher {
	var eventPublisher = new(EventPublisher)
	return eventPublisher
}

func (e *EventPublisher) NewEvent(handlers map[string][]Handler, event Event) {
	e.EventName = event.GetName()
	eventHandlers := handlers[e.EventName]
	for i := range eventHandlers {
		var handler = eventHandlers[i]
		handler.ReactEvent(event)
	}

}
