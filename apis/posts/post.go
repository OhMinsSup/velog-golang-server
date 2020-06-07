package posts

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	post := r.Group("/post")
	{
		post.GET("/:post_id/:url_slug", middlewares.Authorized, controllers.GetPostController)
		post.POST("/", middlewares.Authorized, controllers.WritePostController)
		post.PUT("/:post_id/:url_slug", middlewares.Authorized, controllers.UpdatePostController)
		post.DELETE("/:post_id/:url_slug", middlewares.Authorized, controllers.DeletePostController)
	}
}
