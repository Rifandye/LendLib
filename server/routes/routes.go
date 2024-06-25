package routes

import "github.com/gin-gonic/gin"


func Router(server *gin.Engine) {
	server.POST("/auth/register", register)
	server.POST("/auth/login", login)
	server.POST("/books", addBook)
	server.GET("/books", getAllBooks)
	server.GET("/books/:id", getBookById)
}