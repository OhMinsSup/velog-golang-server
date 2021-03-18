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
		post.PUT("/", middlewares.Authorized, controllers.UpdatePostController)
	}
}
