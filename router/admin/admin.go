package admin

import (
	"gameday/api"
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (a *Admin) InitAdminRouter(Router *gin.RouterGroup) gin.IRoutes {
	adminRouter := Router.Group("admin")
	adminApi := api.ApiGroupApp.Admin
	{
		adminRouter.POST("login", adminApi.UserApi.Login)
	}

	return adminRouter
}
