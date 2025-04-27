package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/models"
)

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	// Bind the incoming JSON to the Event struct
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the event using the Save method
	err := event.CreateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func UpdateEvent(context *gin.Context) {
	// Bind the incoming JSON to the Event struct
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the event using the Save method
	err := event.CreateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func DeleteEvent(context *gin.Context) {
	// Bind the incoming JSON to the Event struct
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the event using the Save method
	err := event.CreateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}
