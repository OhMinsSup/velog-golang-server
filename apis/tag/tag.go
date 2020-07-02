package tag

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	tag := r.Group("/tag")
	{
		tag.GET("/trending", controllers.TrendingTagListController)
	}
}
