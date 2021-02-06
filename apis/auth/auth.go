package auth

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/sendmail", controllers.SendEmailController)
		auth.GET("/code/:code", controllers.CodeController)
		auth.POST("/register/local", controllers.LocalRegisterController)
		auth.POST("/logout", controllers.LogoutController)

		// local
		auth.POST("/register/social", controllers.SocialRegisterController)

		// social
		auth.GET("/social/profile", controllers.SocialProfileController)
		auth.GET("/social/redirect/:provider", controllers.SocialRedirect)
		auth.GET("/social/callback/github", controllers.GithubCallback, controllers.GithubSocialCallback)
		auth.GET("/social/callback/facebook", controllers.FacebookCallback, controllers.FacebookSocialCallback)
	}
}
