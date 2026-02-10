package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "online",
			"pkg":    "api",
		})
	})

	r.Run(":8080")
}
