package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/internal/util"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	claim, err := util.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.Set("claim", claim)
	context.Next()
}
