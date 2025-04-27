package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/db"
	"github.com/hiteshnayak305/go-rest-project/models"
)

func main() {
	db.Initialize()

	// gin.Default() initializes a Gin router with default middleware (logger and recovery)
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	// Render the index.html template
	server.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to the Event Management System",
		})
	})

	// Define a simple GET endpoint
	server.GET("/hello", hello)

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)

	// Start the server on port 8080
	server.Run(":8080")
}

func getEvent(context *gin.Context) {
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

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
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

func hello(context *gin.Context) {
	// Respond with a JSON object
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
