package tests

import (
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
)

func TestAccountHandlerSuccess(t *testing.T) {
	mock := GetMockDB()

	var handler = new(handlers.AccountHandler)
	var eventUserCreated = events.NewEventUserCreated(1)
	var wg sync.WaitGroup

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "accounts" (.+) RETURNING`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "1", "0.00").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	wg.Add(1)
	go handler.HandleEvent(eventUserCreated, &wg)
	wg.Wait()
}
