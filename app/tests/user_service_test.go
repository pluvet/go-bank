package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/services"
)

func TestUserCreateService(t *testing.T) {
	mock := GetMockDB()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" (.+) RETURNING`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	var userExample models.User
	err := services.CreateUser(&userExample)
	assert.Equal(t, userExample.ID, 1)
	assert.NoError(t, err)
}
