package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/routes"
  )

func main() {
	r := gin.Default()
	routes.UserRoute(r)
	r.Run()
}