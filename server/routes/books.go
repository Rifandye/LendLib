package routes

import (
	"net/http"

	"example.com/LendLib/models"
	"github.com/gin-gonic/gin"
)


func addBook(context *gin.Context) {
	var newBook models.Book

	err := context.ShouldBindJSON(&newBook)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = newBook.CreateBook()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save new book", "error": err.Error()})
	}


	context.JSON(http.StatusCreated, gin.H{"message": "Book created successfull", "book": newBook})
}


func getAllBooks(context *gin.Context) {
	books, err := models.GetBooks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch books"})
		return
	}

	context.JSON(http.StatusOK, books)
}