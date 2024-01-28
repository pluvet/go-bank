package services

import (
	"fmt"

	"github.com/pluvet/go-bank/app/eventpublisher"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/repositories"
)

type IAccountService interface {
	CreateAccount(int) (*int, error)
	AccountDeposit(int, float32) (*float32, error)
	AccountWithdraw(int, float32) (*float32, error)
}

type AccountService struct {
	repo           repositories.IAccountRepository
	eventPublisher eventpublisher.EventPublisher
}

func (u *AccountService) CreateAccount(userID int) (*int, error) {

	accountID, err := u.repo.CreateAccount(userID)

	if err != nil {
		return nil, err
	}

	return accountID, nil
}

func NewAccountService(repo repositories.IAccountRepository, eventPublisher *eventpublisher.EventPublisher) AccountService {
	accountService := new(AccountService)
	accountService.repo = repo
	if eventPublisher != nil {
		accountService.eventPublisher = *eventPublisher
	}
	return *accountService
}

func (a *AccountService) AccountDeposit(accountID int, amount float32) (*float32, error) {
	account, findError := a.repo.FindAccount(accountID)
	if findError != nil {
		return nil, findError
	}

	account.Deposit(float32(amount))

	updateError := a.repo.UpdateAccount(account)
	if updateError != nil {
		return nil, updateError
	}

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(amount, account.Balance)
	eventWasPublished := a.eventPublisher.NewEvent(eventAccountBalanceIncreased)
	if !eventWasPublished {
		fmt.Printf("eventAccountBalanceIncreased was not published")
	}
	return &account.Balance, nil
}

func (a *AccountService) AccountWithdraw(accountID int, amount float32) (*float32, error) {
	account, findError := a.repo.FindAccount(accountID)
	if findError != nil {
		return nil, findError
	}

	withdrawError := account.Withdraw(amount)
	if withdrawError != nil {
		return nil, withdrawError
	}

	updateError := a.repo.UpdateAccount(account)
	if updateError != nil {
		return nil, updateError
	}

	var eventAccountBalanceDecreased = events.NewEventAccountBalanceDecreased(amount, account.Balance)
	eventWasPublished := a.eventPublisher.NewEvent(eventAccountBalanceDecreased)
	if !eventWasPublished {
		fmt.Printf("eventAccountBalanceDecreased was not published")
	}
	return &account.Balance, nil
}
