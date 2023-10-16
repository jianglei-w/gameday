package game

import (
	v1 "gameday/api/v1"
	"github.com/gin-gonic/gin"
)

type gameRouter struct {
}

func (gr *gameRouter) InitGameRouters(Router *gin.RouterGroup) (R *gin.IRoutes) {
	gameRouter := Router.Group("admin")
	gameApi := v1.ApiGroupApp.GameApiGroup
	{
		gameRouter.POST("createGame", gameApi.CreateGame)
		gameRouter.GET("lastGame", gameApi.LastGame)
		gameRouter.GET("getallGame", gameApi.GetAllGame)
		gameRouter.GET("getHash", gameApi.GetHash)
	}
	return R
}
