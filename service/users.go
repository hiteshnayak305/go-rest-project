package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshnayak305/go-rest-project/internal/util"
	"github.com/hiteshnayak305/go-rest-project/models"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.CreateUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signup successfully"})
}

func Login(context *gin.Context) {
	var user *models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check password match
	if util.CheckPasswordHash(user.Password, validUser.Password) {
		context.JSON(http.StatusOK, gin.H{"message": "User login successfully"})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User login failed"})
	}
}
