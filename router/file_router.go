package router

import (
	"cloud/controllers/file"
	"cloud/middleware"

	"github.com/gin-gonic/gin"
)

func FileRouter(r *gin.RouterGroup) {
	AuthMiddleware := middleware.AuthMiddleware()
	FileGroup := r.Group("/files")
	{
		FileGroup.GET("/", AuthMiddleware, file.GetFileList)
		FileGroup.POST("/upload", AuthMiddleware, file.UploadFile)
	}

}
