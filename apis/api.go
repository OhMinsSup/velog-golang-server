package apis

import (
	"github.com/OhMinsSup/story-server/apis/auth"
	"github.com/OhMinsSup/story-server/apis/file"
	"github.com/OhMinsSup/story-server/apis/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/v1.0")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		auth.ApplyRoutes(api)
		user.ApplyRoutes(api)
		file.ApplyRoutes(api)
	}
}
