package main

import (
	"github.com/gin-gonic/gin"
)

// Create http server
func main() {
	//Create gin engine
	r := gin.Default()

	//Create a route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Run server
	r.Run(":8080")
}
