package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/models"
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
		case services.ErrorCreatingRecordInDB:
			c.JSON(http.StatusInternalServerError, err.Error())
		case services.ErrorUpdatingRecordInDB:
			c.JSON(http.StatusInternalServerError, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	c.JSON(200, &account)
}

func AccountWithdraw(c *gin.Context) {
	var accountBalanceChange AccountBalanceChange
	c.BindJSON(&accountBalanceChange)
	account, err := services.AccountWithdraw(c.Param("id"), accountBalanceChange.Amount)
	if err != nil {
		switch err.(type) {
		case *services.ErrorFindingOneRecordInDB:
			c.JSON(http.StatusBadRequest, err.Error())
		case *services.ErrorCreatingRecordInDB:
			c.JSON(http.StatusInternalServerError, err.Error())
		case *services.ErrorUpdatingRecordInDB:
			c.JSON(http.StatusInternalServerError, err.Error())
		case *models.WithdrawError:
			c.JSON(http.StatusBadRequest, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	c.JSON(200, &account)

}
