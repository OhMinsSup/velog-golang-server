package apis

import (
	"github.com/OhMinsSup/story-server/apis/auth"
	"github.com/OhMinsSup/story-server/apis/comment"
	"github.com/OhMinsSup/story-server/apis/post"
	"github.com/OhMinsSup/story-server/apis/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/v1.0")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		auth.ApplyRoutes(api)
		post.ApplyRoutes(api)
		posts.ApplyRoutes(api)
		comment.ApplyRoutes(api)
	}
}
