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

	//print log
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	

	//Run server on port 9000
	r.Run(":9000")
}
