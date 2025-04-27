package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(context *gin.Context) {
	// Respond with a JSON object
	context.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
