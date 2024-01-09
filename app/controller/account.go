package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
)

func AccountDeposit(c *gin.Context) {
	var account models.Account
	config.DB.Where("id = ?", c.Param("id")).First(&account)
	amount, err := strconv.ParseFloat(c.Query("amount"), 20)
	if err != nil {
		panic(err)
	}
	account.Deposit(float32(amount))
	c.BindJSON(&account)
	config.DB.Save(&account)

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(float32(amount), account.Balance)
	go publisher.EVENTPUBLISHER.NewEvent(eventAccountBalanceIncreased)

	c.JSON(200, &account)
}

func AccountWithdraw(c *gin.Context) {
	var account models.Account
	config.DB.Where("id = ?", c.Param("id")).First(&account)
	amount, err := strconv.ParseFloat(c.Query("amount"), 20)
	if err != nil {
		panic(err)
	}
	account.Withdraw(float32(amount))
	c.BindJSON(&account)
	config.DB.Save(&account)

	var eventAccountBalanceDecreased = events.NewEventAccountBalanceDecreased(float32(amount), account.Balance)
	go publisher.EVENTPUBLISHER.NewEvent(eventAccountBalanceDecreased)

	c.JSON(200, &account)
}
