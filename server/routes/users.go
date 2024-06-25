package routes

import (
	"net/http"

	"example.com/LendLib/models"
	"example.com/LendLib/utils"
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

func login(context *gin.Context) {
	var user models.UserCredentials

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.FirstName, user.LastName, user.Email, user.Role, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}