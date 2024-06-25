package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin(context *gin.Context) {
	role, exists := context.Get("role")

	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Only admin can do this"})
		return
	}

	if role != "admin" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Only admin can do this"})
		return
	}

	context.Next()
}