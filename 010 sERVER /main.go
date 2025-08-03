package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})

	})

	server.GET("/health", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	server.Run(":8080")

}
