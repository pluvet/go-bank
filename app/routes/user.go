package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/controller"
)

func UserRoute(r *gin.Engine) {
	r.POST("/user", controller.CreateUser)
}
