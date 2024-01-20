package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/models"
	"github.com/pluvet/go-bank/app/services"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := services.CreateUser(&user)
	if err != nil {
		switch err.(type) {
		case *services.ErrorCreatingRecordInDB:
			c.JSON(http.StatusInternalServerError, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	c.JSON(200, &user)
}
