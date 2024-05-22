package routes

import (
	"fmt"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"

	"github.com/gin-gonic/gin"
)

func singnup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, user)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}

	err = user.Login()

	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Printf("Error generating token: %v", err)
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "login successful", "token": token})
}
