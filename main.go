package main

import (
	"example/restapi/db"
	"example/restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
