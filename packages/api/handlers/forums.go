package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GET /forums
func GetForums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of forums"})
}

// POST /forums
func CreateForum(c *gin.Context) {
	c.Status(http.StatusCreated)
}

// GET /forums/:id/threads
func GetThreadsByForum(c *gin.Context) {
	forumID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"forumId": forumID, "threads": []string{}})
}

// POST /forums/:id/threads
func CreateThread(c *gin.Context) {
	c.Status(http.StatusCreated)
}

// GET /threads/:threadId
func GetThreadByID(c *gin.Context) {
	threadID := c.Param("threadId")
	c.JSON(http.StatusOK, gin.H{"threadId": threadID})
}

// DELETE /threads/:threadId
func DeleteThread(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
