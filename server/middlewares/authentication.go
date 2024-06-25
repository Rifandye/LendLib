package middlewares

import (
	"net/http"
	"strings"

	"example.com/LendLib/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(token, bearerPrefix) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
		return
	}

	token = token[len(bearerPrefix):] 

	userId, role, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	context.Set("userId", userId)
	context.Set("role", role)

	context.Next()
}