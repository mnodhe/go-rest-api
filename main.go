package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context) {
	var event = models.Event{}
	err := context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	errr := event.Save()
	if errr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event", "error": errr.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// watched 166
