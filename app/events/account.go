package events

import (
	"sync"

	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

func CreateAccountEvent(userId uint) {
	var wg sync.WaitGroup
	wg.Add(1)
	CreateAccountHandler(&wg, userId)
}

func CreateAccountHandler(wg *sync.WaitGroup, userId uint) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userId)
	config.DB.Create(&account)
	defer wg.Done()
}
