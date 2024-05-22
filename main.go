package main

import (
	"example.com/rest-api/database"
	"example.com/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	database.InitDb()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
