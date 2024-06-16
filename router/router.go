package router

import (
	"cloud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	middleware.InitSessionStore()
	r.Use(middleware.SessionMiddleware())
	api := r.Group("/api")
	UserRouter(api)
	return r
}
