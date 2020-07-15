package file

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	file := r.Group("/file")
	{
		file.POST("/create-url", controllers.CreateUrlController)
		file.POST("/upload", controllers.S3ImageUploadController)
	}
}
