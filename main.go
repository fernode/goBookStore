package main

import (
	"github.com/fernode/goBookStore/config"
	"github.com/fernode/goBookStore/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to our bookstore API! 📚",
		})
	})

	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.AuthRoutes(router, db)

	// Start the server and listen on port 8080
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
