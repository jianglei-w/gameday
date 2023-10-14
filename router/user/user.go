package user

import (
	v1 "gameday/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	userRouter := Router.Group("admin")
	userApi := v1.ApiGroupApp.UserApiGroup
	{
		userRouter.POST("login", userApi.Login)
	}
	return R
}
