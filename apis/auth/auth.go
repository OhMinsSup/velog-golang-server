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
	}
}
