package router

import (
	"cloud/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.RouterGroup) {

	userGroup := c.Group("/")
	{
		userGroup.GET("/checkcode", user.CheckCode)
		//重置密码
		userGroup.POST("/resetpwd", user.ResetPwd)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.Register)
	}
}
