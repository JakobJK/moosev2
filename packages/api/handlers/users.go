package handlers

import (
	"moose-api/db_gen"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Queries *db_gen.Queries
}

// GET /users/:username
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	username := strings.TrimSpace(c.Param("username"))

	user, err := h.Queries.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.Queries.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var params db_gen.CreateUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Queries.CreateUser(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Implementation depends on your sqlc DeleteUser query
	c.Status(http.StatusNoContent)
}

// --- Status & Placeholder Handlers ---

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "online"})
}

