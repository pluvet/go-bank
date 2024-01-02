package events

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

func createAccountEvent(userId int) {
	var wg sync.WaitGroup
	wg.Add(1)
	createAccountHandler(&wg, userId)
}

func createAccountHandler(wg *sync.WaitGroup, c *gin.Context, userId int) {
	var account models.Account
	account.user_id = userId
	c.BindJSON(&account)
	config.DB.Create(&account)
	defer wg.Done()
}
