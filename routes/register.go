package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetFloat64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event", "error": err.Error()})
		return
	}
	isExist, _ := event.IsRegistrationExist(int64(userId))
	if isExist {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user already registered"})
		return
	}

	err = event.Register(int64(userId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event registered", "event": event})
}
func cancelRegistration(context *gin.Context) {
	userId := context.GetFloat64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(int64(userId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel registration", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event cancelled", "event": event})
}
