package main

import (
	"context" 
	"moose-api/database"
	"moose-api/db_gen"
	"moose-api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Capitalize Background
	ctx := context.Background() 
	
	// Capitalize ConnectDB and Close
	pool, err := database.ConnectDB(ctx, "postgres://moose:moose@db:5432/moose_db?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// Capitalize New
	queries := db_gen.New(pool)

	// Capitalize struct names and fields
	userHandler := &handlers.UserHandler{Queries: queries}
	forumHandler := &handlers.ForumHandler{Queries: queries} 

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	    }))
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
		forums.GET("/", forumHandler.GetForums)
		forums.POST("/", forumHandler.CreateForum)
		forums.GET("/:id/threads", forumHandler.GetThreadsByForum)
		forums.POST("/:id/threads", forumHandler.CreateThread)
	}

	r.Run(":8080")
}
