package events

import (
	"reflect"
)

type EventUserCreated struct {
	UserID int
}

func NewEventUserCreated(userID int) *EventUserCreated {
	var eventUserCreated = new(EventUserCreated)
	eventUserCreated.UserID = userID
	return eventUserCreated
}

func (e *EventUserCreated) GetName() string {
	return reflect.TypeOf(e).String()
}

type EventAccountBalanceIncreased struct {
	AddedBalance   float32
	CurrentBalance float32
}

func NewEventAccountBalanceIncreased(addedBalance float32, currentBalance float32) *EventAccountBalanceIncreased {
	var eventAccountBalanceIncreased = new(EventAccountBalanceIncreased)
	eventAccountBalanceIncreased.AddedBalance = addedBalance
	eventAccountBalanceIncreased.CurrentBalance = currentBalance
	return eventAccountBalanceIncreased
}

func (e *EventAccountBalanceIncreased) GetName() string {
	return reflect.TypeOf(e).String()
}
