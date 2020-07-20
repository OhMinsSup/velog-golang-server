package file

import (
	"github.com/OhMinsSup/story-server/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	file := r.Group("/file")
	{
		file.POST("/upload-url", controllers.GeneratePresignedUrlController)
		file.POST("/upload-file", controllers.S3ImageUploadController)
	}
}
