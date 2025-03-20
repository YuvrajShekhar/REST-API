package main

import (
	"example/restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.Use(gin.Logger())
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data", "error": err})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"meesgage": "Event created!", "event": event})

}
