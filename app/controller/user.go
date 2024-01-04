package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/handlers"
	"github.com/pluvet/go-bank/app/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)

	// logic event
	accountHandler := handlers.AccountHandler{}
	var accountEvent = events.NewEvent(&accountHandler, user.ID)
	accountEvent.CreateAccountEvent()

	c.JSON(200, &user)
}
