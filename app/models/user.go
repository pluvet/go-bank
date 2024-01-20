package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Account  *Account
}

type Account struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primary_key"`
	Balance float32 `json:"balance" default:"0"`
	UserID  uint64
	User    User
}

func (a *Account) Deposit(amount float32) {
	a.Balance = a.Balance + amount
}

func (a *Account) Withdraw(amount float32) error {
	if amount > a.Balance {
		return new(WithdrawError)
	}
	a.Balance = a.Balance - amount
	return nil
}

type WithdrawError struct{}

func (w WithdrawError) Error() string {
	return fmt.Sprintf("the requested withdraw amount is bigger than your actual balance")
}
