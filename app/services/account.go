package services

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
)

func AccountDeposit(accountID string, amount float32) (*models.Account, error) {
	var account models.Account
	config.DB.Where("id = ?", accountID).First(&account)
	if account == (models.Account{}) {
		err := new(ErrorFindingOneRecordInDB)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	account.Deposit(float32(amount))
	result := config.DB.Save(&account)
	if result.Error != nil {
		err := new(ErrorUpdatingRecordInDB)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(amount, account.Balance)
	eventPublisher := publisher.GetEventPublisher()
	go eventPublisher.NewEvent(eventAccountBalanceIncreased)

	return &account, nil
}

func AccountWithdraw(accountID string, amount float32) (*models.Account, error) {
	var account models.Account
	config.DB.Where("id = ?", accountID).First(&account)
	if account == (models.Account{}) {
		err := new(ErrorFindingOneRecordInDB)
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
		err := new(ErrorUpdatingRecordInDB)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}

	var eventAccountBalanceDecreased = events.NewEventAccountBalanceDecreased(amount, account.Balance)
	eventPublisher := publisher.GetEventPublisher()
	go eventPublisher.NewEvent(eventAccountBalanceDecreased)

	return &account, nil
}
