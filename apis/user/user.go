package user

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/:username", controllers.GetUserProfile)
	}
}
