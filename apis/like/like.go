package like

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	like := r.Group("/like")
	{
		like.POST("/:post_id", middlewares.Authorized, controllers.LikePostController)
		like.DELETE("/:post_id", middlewares.Authorized, controllers.UnLikePostController)
	}
}
