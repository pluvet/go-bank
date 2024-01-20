package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/services"
)

func TestAccountDepositServiceSuccess(t *testing.T) {
	mock := GetMockDB()

	rows := sqlmock.NewRows([]string{"id", "balance", "user_id"}).AddRow(1, 0, 1)

	mock.ExpectQuery("SELECT (.+) FROM \"accounts\" WHERE id =(.+)").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"accounts\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	account, err := services.AccountDeposit("1", 25.2)
	assert.Equal(t, account.Balance, float32(25.2))
	assert.NoError(t, err)
}

func TestAccountWithdrawServiceSuccess(t *testing.T) {
	mock := GetMockDB()

	rows := sqlmock.NewRows([]string{"id", "balance", "user_id"}).AddRow(1, 50, 1)

	mock.ExpectQuery("SELECT (.+) FROM \"accounts\" WHERE id =(.+)").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"accounts\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	account, err := services.AccountWithdraw("1", 25)
	assert.Equal(t, account.Balance, float32(25))
	assert.NoError(t, err)
}

func TestAccountWithdrawServiceFail(t *testing.T) {
	//withdraw should fail when balance is less than withdraw amount
	mock := GetMockDB()

	rows := sqlmock.NewRows([]string{"id", "balance", "user_id"}).AddRow(1, 0, 1)

	mock.ExpectQuery("SELECT (.+) FROM \"accounts\" WHERE id =(.+)").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"accounts\" SET .+").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	account, err := services.AccountWithdraw("1", 25)
	assert.Nil(t, account)
	testError := new(models.WithdrawError)
	assert.EqualError(t, err, testError.Error())
}
