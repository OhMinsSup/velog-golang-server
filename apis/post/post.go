package post

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	post := r.Group("/post")
	{
		post.POST("/", middlewares.Authorized, controllers.WritePostController)

		post.GET("/:post_id/", middlewares.Authorized, controllers.GetPostController)
		post.PUT("/:post_id/", middlewares.Authorized, controllers.UpdatePostController)
		post.DELETE("/:post_id/", middlewares.Authorized, controllers.DeletePostController)

		post.POST("/:post_id/like", middlewares.Authorized, controllers.LikePostController)
		post.DELETE("/:post_id/like", middlewares.Authorized, controllers.UnLikePostController)
	}
}
