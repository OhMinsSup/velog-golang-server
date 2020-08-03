package user

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/OhMinsSup/story-server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/", middlewares.Authorized, controllers.GetCurrentUserController)
		user.GET("/:username", controllers.GetUserProfileController)

		user.PUT("/email-rules", middlewares.Authorized, controllers.UpdateEmailRulesController)
		user.PUT("/social-info", middlewares.Authorized, controllers.UpdateSocialController)
		user.PUT("/profile", middlewares.Authorized, controllers.UpdateProfileController)
	}
}
