package main

import (
	"example.com/project/db"
	"example.com/project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//init db
	db.InitDB()
	//set up enginecome with logger and recovery middleware (can use http server)
	server := gin.Default()

	routes.RegisterRoutes(server)

	//start server on localhost:8080
	server.Run(":8080")
}

// func getEvents(context *gin.Context) {
// 	events, err := models.GetAllEvents()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events. Try again later."})
// 		return
// 	}

// 	//create response and return in JSON format
// 	//slice and map are common type of data that will be sent back
// 	//http.statusOK = status code 200, gin.H is custom type of map that provide by gin
// 	context.JSON(http.StatusOK, gin.H{"message": events})
// }

// func getEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

// 	if err != nil {
// 		fmt.Println(err)
// 		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse event id."})
// 		return
// 	}
// 	event, err := models.GetEventById(eventId)
// 	if err != nil {
// 		fmt.Println(err)
// 		context.JSON(http.StatusInternalServerError,gin.H{"message": "Could not fetch event."})
// 		return
// 	}
// 	context.JSON(http.StatusOK, event)
// }

// func createEvents(context *gin.Context) {
// 	var event models.Event

// 	//ShouldBindJson func will look to the request body and store that in the variable
// 	// client must sent data in shape of variable type but some field of variable can be null
// 	err := context.ShouldBindJSON(&event)
// 	if err != nil {
// 		fmt.Println(err)
// 		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse request data."})
// 		return
// 	}

// 	//save request
// 	err = event.Save()
// 	if err != nil {
// 		fmt.Println(err)
// 		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events. Try again later."})
// 		return
// 	}
// 	context.JSON(http.StatusCreated,gin.H{"message": "Event created!", "event":event})
// }