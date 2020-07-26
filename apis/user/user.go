package user

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/", middlewares.Authorized, controllers.GetCurrentUser)
		user.GET("/:username", controllers.GetUserProfile)
	}
}
