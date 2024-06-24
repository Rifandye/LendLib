package main

import (
	"example.com/LendLib/db"
	"example.com/LendLib/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.Router(server)

	server.Run(":8080")
}