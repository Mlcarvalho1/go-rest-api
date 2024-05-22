package routes

import (
	"strconv"

	"example.com/rest-api/models"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(500, gin.H{"error": "could not fetch event"})
		return
	}

	context.JSON(200, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(500, gin.H{"error": "could not fetch event"})
		return
	}

	if event.UserID != userId {
		context.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = id
	err = updateEvent.Update()

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, updateEvent)
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(500, gin.H{"error": "could not fetch event"})
		return
	}

	if event.UserID != userId {
		context.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(500, gin.H{"error": "could not delete event"})
		return
	}

	context.JSON(200, gin.H{"message": "event deleted"})
}
