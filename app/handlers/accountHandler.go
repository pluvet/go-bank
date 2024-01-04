package handlers

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

type AccountHandler struct {
	handler Handler
}

func (a *AccountHandler) AccountCreate(userId uint) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userId)
	config.DB.Create(&account)
}
