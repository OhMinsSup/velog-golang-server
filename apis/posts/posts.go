package posts

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.GET("/", controllers.ListPostsController)
		posts.GET("/trending", controllers.TrendingPostsController)
		posts.GET("/reading", controllers.ReadingPostsController)
		posts.GET("/likes", controllers.LikePostsController)
	}
}
