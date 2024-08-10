package routes

import (
	"example.com/project/middlewares"
	"github.com/gin-gonic/gin"
)

//registering all routes
//get to the existing server by passing the pointer of gin engine as argument
func RegisterRoutes(server *gin.Engine) {
	//when request come to the server gin will pass context as a parameter to the 
	//eventhandler function automatically
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	//protected routes
	// server.POST("/events", middlewares.Authenticate,createEvent)
	// server.PUT("/events/:id", middlewares.Authenticate,editEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate,deleteEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", editEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}