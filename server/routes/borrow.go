package routes

import (
	"net/http"
	"strconv"
	"strings"

	"example.com/LendLib/models"
	"example.com/LendLib/utils"
	"github.com/gin-gonic/gin"
)

func borrowBook(context *gin.Context) {
	//authentication for bearer token jwt
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(token, bearerPrefix) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
		return
	}

	token = token[len(bearerPrefix):] 

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	
	//borrowing logic
	var newBorrow models.Borrow

	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id", "error": err.Error()})
		return
	}

	err = newBorrow.CreateBorrow(bookId, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not borrow"})
		return
	}


	context.JSON(http.StatusCreated, gin.H{"message": "Borrow created successfully"})
}