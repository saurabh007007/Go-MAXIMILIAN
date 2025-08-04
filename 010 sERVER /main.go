package main

import (
	"net/http"

	"http.com/com/db"

	"github.com/gin-gonic/gin"
	"http.com/com/models"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})

	})

	server.GET("/event", GetEvenets)
	server.POST("/event", CreateEvents)

	server.GET("/health", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	server.Run(":8080")

}
func GetEvenets(res *gin.Context) {
	events := models.GetAllEvents()
	res.JSON(http.StatusAccepted, events)
}

func CreateEvents(res *gin.Context) {
	var event models.Event
	err := res.ShouldBindJSON(&event)

	if err != nil {
		res.JSON(http.StatusBadRequest, gin.H{
			"message": "Something is missing ",
		})

	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	res.JSON(http.StatusCreated, gin.H{
		"messege": "Evenet created",
		"event":   event,
	})

}
