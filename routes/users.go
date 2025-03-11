package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}
	email, err := user.IsEmailExist(user.Email)
	if err != nil {
		return
	}
	if email {
		context.JSON(http.StatusBadRequest, gin.H{"message": "email address is already exist."})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save user data.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials.", "error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful."})
}
