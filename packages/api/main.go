package main

import (
	"moose-api/database"
	"moose-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize DB and get sqlc queries
	queries, dbConn := database.Init()
	defer dbConn.Close()

	// 2. Inject queries into handlers
	userHandler := &handlers.UserHandler{Queries: queries}

	r := gin.Default()

	// --- ROUTES ---
	r.GET("/status", handlers.GetStatus)

	users := r.Group("/users")
	{
		users.GET("/", userHandler.ListUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:username", userHandler.GetUserProfile)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	forums := r.Group("/forums")
	{
		forums.GET("/", handlers.GetForums)
		forums.POST("/", handlers.CreateForum)
		forums.GET("/:id/threads", handlers.GetThreadsByForum)
		forums.POST("/:id/threads", handlers.CreateThread)
	}

	r.Run(":8080")
}
