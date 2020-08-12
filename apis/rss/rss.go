package rss

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	rss := r.Group("/rss")
	{
		rss.GET("/", controllers.GetEntireFeed)
	}
}
