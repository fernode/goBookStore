package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin router
	router := gin.Default()

	// Define a basic route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to our bookstore API! 📚",
		})
	})

	// Start the server and listen on port 8080
	router.Run(":8080")
}
