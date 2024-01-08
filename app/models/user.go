package models

import "gorm.io/gorm"

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

func (a *Account) Withdraw(amount float32) {
	if amount <= a.Balance {
		a.Balance = a.Balance - amount
	} else {

	}

}
