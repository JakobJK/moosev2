package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moose-api/db_gen"
)

type ForumHandler struct {
	Queries *db_gen.Queries
}

// GET /forums
func (h *ForumHandler) GetForums(c *gin.Context) {
	forums, err := h.Queries.ListForums(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forums"})
		return
	}

	c.JSON(http.StatusOK, forums)
}

// POST /forums
func (h *ForumHandler) CreateForum(c *gin.Context) {
	// Logic for creating a forum
	c.Status(http.StatusCreated)
}

// GET /forums/:id/threads
func (h *ForumHandler) GetThreadsByForum(c *gin.Context) {
	forumID := c.Param("id")
	// Future: use h.Queries.ListThreadsByForum
	c.JSON(http.StatusOK, gin.H{"forumId": forumID, "threads": []string{}})
}

// POST /forums/:id/threads
func (h *ForumHandler) CreateThread(c *gin.Context) {
	// Logic for creating a thread
	c.Status(http.StatusCreated)
}

// GET /threads/:threadId
func (h *ForumHandler) GetThreadByID(c *gin.Context) {
	threadID := c.Param("threadId")
	c.JSON(http.StatusOK, gin.H{"threadId": threadID})
}

// DELETE /threads/:threadId
func (h *ForumHandler) DeleteThread(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
