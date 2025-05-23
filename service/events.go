package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/internal/util"
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

	claim, err := util.GetCustomClaim(context, "claim")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid claim type"})
		return
	}
	event.UserID = claim.UserID

	// Save the event using the Save method
	err = event.CreateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func UpdateEvent(context *gin.Context) {
	// Bind id from param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Bind the incoming JSON to the Event struct
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetch existing event
	_, err = models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claim, err := util.GetCustomClaim(context, "claim")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid claim type"})
		return
	}

	if event.UserID != claim.UserID {
		context.JSON(http.StatusForbidden, gin.H{"error": "User not authorized to operate on this object"})
		return
	}

	// updated event
	event.ID = id
	event.UpdateEvent()

	// Respond with a success message
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	// Bind id from param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetch existing event
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claim, err := util.GetCustomClaim(context, "claim")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid claim type"})
		return
	}

	if event.UserID != claim.UserID {
		context.JSON(http.StatusForbidden, gin.H{"error": "User not authorized to operate on this object"})
		return
	}

	// Delete event
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusNoContent, gin.H{"message": "Event deleted successfully"})
}
