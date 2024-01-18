package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
	"github.com/pluvet/go-bank/app/services"
)

type AccountBalanceChange struct {
	Amount float32 `json:"amount"`
}

func AccountDeposit(c *gin.Context) {
	var accountBalanceChange AccountBalanceChange
	c.BindJSON(&accountBalanceChange)
	account, err := services.AccountDeposit(c.Param("id"), accountBalanceChange.Amount)
	if err != nil {
		switch err.(type) {
		case services.CreateError:
			c.JSON(http.StatusInternalServerError, err.Error())
		case services.UpdateError:
			c.JSON(http.StatusInternalServerError, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	var eventAccountBalanceIncreased = events.NewEventAccountBalanceIncreased(accountBalanceChange.Amount, account.Balance)
	go publisher.EventPublisher.NewEvent(eventAccountBalanceIncreased)

	c.JSON(200, &account)
}

func AccountWithdraw(c *gin.Context) {
	var accountBalanceChange AccountBalanceChange
	c.BindJSON(&accountBalanceChange)
	account, err := services.AccountWithdraw(c.Param("id"), accountBalanceChange.Amount)
	if err != nil {
		println(err.Error())
		switch err.(type) {
		case *services.FindError:
			c.JSON(http.StatusBadRequest, err.Error())
		case *services.CreateError:
			c.JSON(http.StatusInternalServerError, err.Error())
		case *services.UpdateError:
			c.JSON(http.StatusInternalServerError, err.Error())
		case *models.WithdrawError:
			println("error")
			c.JSON(http.StatusBadRequest, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	} else {
		var eventAccountBalanceDecreased = events.NewEventAccountBalanceDecreased(accountBalanceChange.Amount, account.Balance)
		go publisher.EventPublisher.NewEvent(eventAccountBalanceDecreased)

		c.JSON(200, &account)
	}

}
