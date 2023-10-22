package admin

import (
	v1 "gameday/api/v1"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

// InitAdminRouter 初始化管理员登录路由，无鉴权
func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	userRouter := Router.Group("admin")
	userApi := v1.ApiGroupApp.AdminApiGroup
	{
		userRouter.POST("login", userApi.Login)
	}
	return R
}
