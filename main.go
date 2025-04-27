package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/db"
	"github.com/hiteshnayak305/go-rest-project/routes"
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

	routes.RegisterRoutes(server)

	// Start the server on port 8080
	server.Run(":8080")
}
