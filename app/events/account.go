package events

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

func CreateAccountEvent(userId uint) {
	go CreateAccountHandler(userId)
}

func CreateAccountHandler(userId uint) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userId)
	config.DB.Create(&account)
}
