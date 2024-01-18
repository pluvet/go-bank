package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/events"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/publisher"
	"github.com/pluvet/go-bank/app/services"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := services.CreateUser(&user)
	if err != nil {
		switch err.(type) {
		case *services.CreateError:
			c.JSON(http.StatusInternalServerError, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	var eventUserCreated = events.NewEventUserCreated(user.ID)
	go publisher.EventPublisher.NewEvent(eventUserCreated)

	c.JSON(200, &user)
}
