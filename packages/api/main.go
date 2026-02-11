package main

import (
	"moose-api/handlers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()


	// --- HEALTH CHECK ---
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "online"})
	})

	// --- USER ROUTES ---
	users := r.Group("/users")
	{
		users.GET("/", getUsers)          // List users (Paginated)
		users.POST("/", createUser)       // Register
		users.GET("/:username", handlers.GetUserProfile)
		users.DELETE("/:id", deleteUser)  // Delete
	}

	// --- FORUM & THREAD ROUTES ---
	forums := r.Group("/forums")
	{
		forums.GET("/", getForums)
		forums.POST("/", createForum)

		// Nested: Threads belong to a forum
		forums.GET("/:id/threads", getThreadsByForum)
		forums.POST("/:id/threads", createThread)
	}

	// Individual thread actions
	threads := r.Group("/threads")
	{
		threads.GET("/:threadId", getThreadByID)
		threads.DELETE("/:threadId", deleteThread)
	}

	r.Run(":8080")
}


// Placeholder handlers so the code compiles
func getThreadsByForum(c *gin.Context){ c.Status(200) }
func getUsers(c *gin.Context)         { c.Status(200) }
func createUser(c *gin.Context)       { c.Status(201) }
func getUserByID(c *gin.Context)      { c.Status(200) }
func deleteUser(c *gin.Context)       { c.Status(204) }
func getForums(c *gin.Context)        { c.Status(200) }
func createForum(c *gin.Context)      { c.Status(201) }
func createThread(c *gin.Context)     { c.Status(201) }
func getThreadByID(c *gin.Context)    { c.Status(200) }
func deleteThread(c *gin.Context)     { c.Status(204) }
