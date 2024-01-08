package events

import (
	"reflect"
)

type EventUserCreated struct {
	UserID int
}

func NewEventUserCreated(userID int) *EventUserCreated {
	var a = new(EventUserCreated)
	a.UserID = userID
	return a
}

func (e *EventUserCreated) GetName() string {
	return reflect.TypeOf(e).String()
}
