package routes

import (
	"net/http"

	"example.com/LendLib/models"
	"github.com/gin-gonic/gin"
)


func register(context *gin.Context) {
	var newUser models.User

	err := context.ShouldBindJSON(&newUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = newUser.CreateUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	newUser.Password = ""

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}