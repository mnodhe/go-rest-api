package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messsage": "invalid id"})
		return
	}
	eventModel, errr := models.GetEventById(eventId)
	if errr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messsage": "invalid id"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": eventModel})

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
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}
	var updatedEvent = models.Event{}
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse req data", "error": err.Error()})
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse req data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	result, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}
	err = result.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted", "event": result})

}
