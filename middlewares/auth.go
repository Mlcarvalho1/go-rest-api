package middlewares

import (
	"fmt"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"error": "authorization token not provided"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		fmt.Printf("Error validating token: %v", err)
		context.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
