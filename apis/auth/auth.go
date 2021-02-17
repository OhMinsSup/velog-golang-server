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
		auth.GET("/social/redirect/:provider", controllers.SocialRedirectController)
		auth.GET("/social/callback/github", controllers.SocialGithubCallbackController, controllers.SocialCallbackController)
		auth.GET("/social/callback/facebook", controllers.SocialFacebookCallbackController, controllers.SocialCallbackController)
		auth.GET("/social/callback/kakao", controllers.SocialKakaoCallbackController, controllers.SocialCallbackController)


		// local
		//auth.POST("/register/social", controllers.SocialRegisterController)

		// social
		//auth.GET("/social/profile", controllers.SocialProfileController)
	}
}
