package routes

import (
	"net/http"
	"strconv"

	"example.com/LendLib/models"
	"github.com/gin-gonic/gin"
)

func borrowBook(context *gin.Context) {
	var newBorrow models.Borrow

	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "details": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	err = newBorrow.CreateBorrow(bookId, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not borrow", "details": err.Error()})
		return
	}


	context.JSON(http.StatusCreated, gin.H{"message": "Borrow created successfully"})
}

func returnBook(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "details": err.Error()})
		return
	}

	err = models.ReturnBook(bookId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not return book", "details": err.Error()})
		return
	}


	context.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}

func getBorrowedBook(context *gin.Context) {
	userId := context.GetInt64("userId")

	userWithBooks, err := models.GetUserBorrowedBooks(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving borrowed books", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userWithBooks)
}