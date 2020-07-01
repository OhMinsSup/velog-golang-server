package posts

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.GET("/", controllers.ListPostsController)
		posts.GET("/trending", controllers.TrendingPostsController)
		posts.GET("/reading", middlewares.Authorized, controllers.ReadingPostsController)
		posts.GET("/likes", middlewares.Authorized, controllers.LikePostsController)
	}
}
