package comment

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	comment := r.Group("/comment")
	{
		comment.POST("/", controllers.WriteCommentController)
	}
}
