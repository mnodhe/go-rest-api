package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
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
