package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/internal/util"
	"github.com/hiteshnayak305/go-rest-project/models"
)

func RegisterForEvent(context *gin.Context) {
	claim, err := util.GetCustomClaim(context, "claim")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid claim type"})
		return
	}

	// Bind id from param
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = event.Register(claim.UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func CancelRegisterForEvent(context *gin.Context) {
	claim, err := util.GetCustomClaim(context, "claim")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid claim type"})
		return
	}

	// Bind id from param
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = event.CancelRegistration(claim.UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User un-registered successfully"})
}
