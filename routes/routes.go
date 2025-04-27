package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/internal/middlewares"
	"github.com/hiteshnayak305/go-rest-project/service"
)

func RegisterRoutes(server *gin.Engine) {
	// Define a simple GET endpoint
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/healthz", service.Healthz)

	server.GET("/events", service.GetEvents)
	server.GET("/events/:id", service.GetEvent)

	authenticated.POST("/events", service.CreateEvent)
	authenticated.PUT("/events/:id", service.UpdateEvent)
	authenticated.DELETE("/events/:id", service.DeleteEvent)

	server.POST("/signup", service.SignUp)
	server.POST("/login", service.Login)

}
