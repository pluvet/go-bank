package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/routes"
  )

func main() {
	r := gin.Default()
	routes.UserRoute(r)
	r.Run()
}