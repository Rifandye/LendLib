package routes

import (
	"example.com/LendLib/middlewares"
	"github.com/gin-gonic/gin"
)


func Router(server *gin.Engine) {
	server.POST("/auth/register", register)
	server.POST("/auth/login", login)

	admin := server.Group("/")
	admin.Use(middlewares.Authenticate)
	admin.Use(middlewares.AuthorizeAdmin)
	admin.POST("/books", addBook)


	server.GET("/books", getAllBooks)
	server.GET("/books/:id", getBookById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/borrow/:id", borrowBook)
	authenticated.PUT("/borrow/:id", returnBook)
}