package user

import (
	v1 "gameday/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	router := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApiGroup
	{
		router.POST("login", userApi.Login)
	}

	return R
}

func (ur *UserRouter) InitAuthRouters(Router *gin.RouterGroup) (R *gin.IRoutes) {
	router := Router.Group("user")
	gameApi := v1.ApiGroupApp.GameApiGroup
	userApi := v1.ApiGroupApp.UserApiGroup
	{
		router.POST("userSubject", gameApi.GetQuestionsByGameId)
		router.POST("updateName", userApi.UpdateUsername)
		router.POST("getUser", userApi.GetUserByHash)
		router.POST("completeID", userApi.GetSuccessQuestionByUserId)
		router.POST("usergrade", userApi.UserScore)
	}

	return R
}
