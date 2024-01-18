package services

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

func AccountDeposit(accountID string, amount float32) (*models.Account, error) {
	var account models.Account
	config.DB.Where("id = ?", accountID).First(&account)
	if account == (models.Account{}) {
		err := new(FindError)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	account.Deposit(float32(amount))
	result := config.DB.Save(&account)
	if result.Error != nil {
		err := new(UpdateError)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	return &account, nil
}

func AccountWithdraw(accountID string, amount float32) (*models.Account, error) {
	var account models.Account
	config.DB.Where("id = ?", accountID).First(&account)
	if account == (models.Account{}) {
		err := new(FindError)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	err := account.Withdraw(amount)
	if err != nil {
		return nil, err
	}
	result := config.DB.Save(&account)
	if result.Error != nil {
		err := new(UpdateError)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	return &account, nil
}
