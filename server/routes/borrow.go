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
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	err = newBorrow.CreateBorrow(bookId, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not borrow"})
		return
	}


	context.JSON(http.StatusCreated, gin.H{"message": "Borrow created successfully"})
}