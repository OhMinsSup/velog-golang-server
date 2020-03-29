package auth

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		// local
		auth.POST("/register/local", controllers.LocalRegisterController)
		auth.POST("/sendmail", controllers.SendEmailController)
		auth.GET("/code/:code", controllers.CodeController)
		// social

		auth.GET("/social/redirect/:provider", controllers.SocialRedirect)
		auth.GET("/callback/github", controllers.GithubCallback, controllers.SocialCallback)
		auth.GET("/callback/google", controllers.GoogleCallback, controllers.SocialCallback)
		auth.GET("/callback/facebook", controllers.FacebookCallback, controllers.SocialCallback)
	}
}
