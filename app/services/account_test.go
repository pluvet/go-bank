package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/repositories"
)

func TestCreateAccountService(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := repositories.NewMockIAccountRepository(ctrl)

	id := 1
	m.
		EXPECT().
		CreateAccount(1).
		Return(&id, nil).
		Times(1)
	eventpublisher := GetMockEventPublisher()
	accountService := NewAccountService(m, eventpublisher)
	accountID, err := accountService.CreateAccount(1)
	assert.Equal(t, &id, accountID)
	assert.NoError(t, err)
}

func TestAccountDepositServiceSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := repositories.NewMockIAccountRepository(ctrl)

	var account models.Account
	account.ID = 1
	account.Balance = 0
	m.
		EXPECT().
		FindAccount(1).
		Return(&account, nil).
		Times(1)

	m.
		EXPECT().
		UpdateAccount(&account).
		Return(nil).
		Times(1)

	accountService := NewAccountService(m, GetMockEventPublisher())
	newBalance, err := accountService.AccountDeposit(1, 12.3)
	assert.Equal(t, float32(12.3), *newBalance)
	assert.NoError(t, err)
}

func TestAccountWithdrawServiceSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := repositories.NewMockIAccountRepository(ctrl)

	var account models.Account
	account.ID = 1
	account.Balance = 50
	m.
		EXPECT().
		FindAccount(1).
		Return(&account, nil).
		Times(1)

	m.
		EXPECT().
		UpdateAccount(&account).
		Return(nil).
		Times(1)

	accountService := NewAccountService(m, GetMockEventPublisher())
	newBalance, err := accountService.AccountWithdraw(1, 10.0)
	assert.Equal(t, float32(40.0), *newBalance)
	assert.NoError(t, err)
}

func TestAccountWithdrawServiceFail(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := repositories.NewMockIAccountRepository(ctrl)

	var account models.Account
	account.ID = 1
	account.Balance = 0
	m.
		EXPECT().
		FindAccount(1).
		Return(&account, nil).
		Times(1)

	m.
		EXPECT().
		UpdateAccount(&account).
		Return(nil).
		Times(0)

	accountService := NewAccountService(m, GetMockEventPublisher())
	newBalance, err := accountService.AccountWithdraw(1, 10.0)
	assert.Nil(t, newBalance)
	assert.Error(t, models.WithdrawError{}, err.Error())
}
