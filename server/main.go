package main

import (
	"example.com/LendLib/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.Router(server)

	server.Run(":8080")
}