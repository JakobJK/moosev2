package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetForums(c *gin.Context) {
	c.JSON(200, gin.H{"message": "forums placeholder"})
}
