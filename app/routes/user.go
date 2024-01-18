package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/controllers"
)

func UserRoute(r *gin.Engine) {
	r.POST("/user", controllers.CreateUser)
	r.POST("/account/:id/deposit", controllers.AccountDeposit)
	r.POST("/account/:id/withdraw", controllers.AccountWithdraw)
}
