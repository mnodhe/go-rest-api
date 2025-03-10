package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/db"
	"go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

// watched 166
