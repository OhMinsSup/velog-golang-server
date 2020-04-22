package apis

import (
	"github.com/OhMinsSup/story-server/apis/auth"
	"github.com/OhMinsSup/story-server/apis/posts"
	"github.com/gin-gonic/gin"
	"log"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/v1.0")
	{
		api.GET("/health", func(c *gin.Context) {
			log.Println(c.Cookie("access_token"))
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		auth.ApplyRoutes(api)
		posts.ApplyRoutes(api)
	}
}
