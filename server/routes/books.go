package routes

import (
	"net/http"
	"strconv"

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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save new book", "details": err.Error()})
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

func getBookById(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "details": err.Error()})
			return
	}

	book, err := models.GetBook(bookId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch book", "details": err.Error()})
		return
	}


	context.JSON(http.StatusOK, book)
}

func updateBookById(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "details": err.Error()})
			return
	}

	_, err = models.GetBook(bookId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch book", "details": err.Error()})
		return
	}

	var updatedBook models.Book

	err = context.ShouldBindJSON(&updatedBook)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	updatedBook.ID = bookId

	err = updatedBook.UpdateBook()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update book", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book update successfully"})
}

func deleteBookById(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "details": err.Error()})
			return
	}

	book, err := models.GetBook(bookId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch book", "details": err.Error()})
		return
	}


	err = book.DeleteBook()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete book", "details": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}