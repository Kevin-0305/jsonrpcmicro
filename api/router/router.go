package router

import (
	auth "jsonrpcmicro/api/auth"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	ProjectRouter := Router.Group("auth")
	{
		ProjectRouter.POST("Login", auth.Login) // 新建Project
	}
}
