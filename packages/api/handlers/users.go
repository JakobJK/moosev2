package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserProfile struct {
	Username       string   `json:"username"`
	Bio            string   `json:"bio"`
	AvatarURL      *string  `json:"avatarUrl"`
	JoinedDate     string   `json:"joinedDate"`
	Posts          int      `json:"posts"`      // Moved to top level
	Topics         int      `json:"topics"`     // Moved to top level
	Reputation     int      `json:"reputation"` // Moved to top level
	RecentActivity []string `json:"recentActivity"`
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")

	profile := UserProfile{
		Username:       username,
		Bio:            "3D Environment Artist specializing in modular workflows and trim sheets.",
		AvatarURL:      nil,
		JoinedDate:     "2024-11-12",
		Posts:          154,
		Topics:         12,
		Reputation:     850,
		RecentActivity: []string{
			"Low Poly Sci-Fi Door (WIP) - 2 hours ago",
			"Modular Building Kit (Showcase) - 3 days ago",
		},
	}

	c.JSON(http.StatusOK, profile)
}
