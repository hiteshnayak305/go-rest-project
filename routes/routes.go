package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/service"
)

func RegisterRoutes(server *gin.Engine) {
	// Define a simple GET endpoint
	server.GET("/healthz", service.Healthz)

	server.GET("/events", service.GetEvents)
	server.POST("/events", service.CreateEvent)
	server.GET("/events/:id", service.GetEvent)
	server.PUT("/events/:id", service.UpdateEvent)
	server.DELETE("/events/:id", service.DeleteEvent)
}
