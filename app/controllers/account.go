package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
	"github.com/pluvet/go-bank/app/repositories"
	"github.com/pluvet/go-bank/app/services"
)

type AccountBalanceChangeInputDTO struct {
	Amount float32
}

type AccountBalanceChangeOutputDTO struct {
	Balance float32
}

func AccountDeposit(c *gin.Context) {
	var accountBalanceChangeDTO AccountBalanceChangeInputDTO
	c.BindJSON(&accountBalanceChangeDTO)
	stringID := c.Param("id")
	accountID, idError := strconv.Atoi(stringID)
	if idError != nil {
		c.JSON(400, "id is not an int value")
	}
	accountService := services.NewAccountService(new(repositories.AccountRepository), publisher.GetEventPublisher())
	newBalance, err := accountService.AccountDeposit(accountID, accountBalanceChangeDTO.Amount)
	if err != nil {
		switch err.(type) {
		case *repositories.ErrorFindingOneRecordInDB:
			c.JSON(404, err.Error())
		case *repositories.ErrorUpdatingRecordInDB:
			c.JSON(500, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	var output AccountBalanceChangeOutputDTO
	output.Balance = *newBalance
	c.JSON(200, output)
}

func AccountWithdraw(c *gin.Context) {
	var accountBalanceChangeDTO AccountBalanceChangeInputDTO
	c.BindJSON(&accountBalanceChangeDTO)
	stringID := c.Param("id")
	accountID, idError := strconv.Atoi(stringID)
	if idError != nil {
		c.JSON(400, "id is not an int value")
	}
	accountService := services.NewAccountService(new(repositories.AccountRepository), publisher.GetEventPublisher())
	newBalance, err := accountService.AccountWithdraw(accountID, accountBalanceChangeDTO.Amount)
	if err != nil {
		switch err.(type) {
		case *repositories.ErrorFindingOneRecordInDB:
			c.JSON(404, err.Error())
		case *repositories.ErrorUpdatingRecordInDB:
			c.JSON(500, err.Error())
		case *models.WithdrawError:
			c.JSON(400, err.Error())
		default:
			c.JSON(500, "Internal Server Error")
		}
	}
	var output AccountBalanceChangeOutputDTO
	output.Balance = *newBalance
	c.JSON(200, output)
}
