package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(500, gin.H{"error": "could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "registration successful"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(500, gin.H{"error": "could not fetch event"})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "registration cancelled"})
}
