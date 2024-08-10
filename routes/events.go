package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/project/models"
	"github.com/gin-gonic/gin"
)

// pass context as pointer to make sure that we work on one single context object
// and must be used in order to send back a response to the originally received request
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events. Try again later."})
		return
	}

	//create response and return in JSON format
	//slice and map are common type of data that will be sent back
	//http.statusOK = status code 200, gin.H is custom type of map that provide by gin
	context.JSON(http.StatusOK, gin.H{"message": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event

	//ShouldBindJson func will look to the request body and store that in the variable
	// client must sent data in shape of variable type but some field of variable can be null
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse request data."})
		return
	}
	event.UserID = context.GetInt64("userId")
	//save request
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message": "Event created!", "event":event})
}

func editEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not fetch event."})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse request data."})
		return
	}
	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized,gin.H{"message": "not authorized to update event."})
		return
	}
	//assign the existing id that we want to edit the data
	updatedEvent.ID = eventId
	err = updatedEvent.Updated()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not updated event."})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message": "Event updated!", "event":updatedEvent})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not fetch event."})
		return
	}
	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized,gin.H{"message": "not authorized to delete event."})
		return
	}
	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not delete event."})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": "Event deleted!"})
}