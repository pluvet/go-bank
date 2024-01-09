package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)

	var eventUserCreated = events.NewEventUserCreated(user.ID)
	go publisher.EVENTPUBLISHER.NewEvent(eventUserCreated)

	c.JSON(200, &user)
}
