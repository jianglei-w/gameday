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
	questApi := v1.ApiGroupApp.QuestApiGroup
	{
		gameRouter.POST("createGame", gameApi.CreateGame)
		gameRouter.GET("lastGame", gameApi.LastGame)
		gameRouter.GET("getallGame", gameApi.GetAllGame)
		gameRouter.GET("getHash", gameApi.GetHash)
		gameRouter.GET("RankList", gameApi.RankList)
		gameRouter.POST("startGame", gameApi.StartGame)
		gameRouter.POST("stopGame", gameApi.StopGame)
		gameRouter.POST("deleteGame", gameApi.DoneGame)
		gameRouter.POST("success", gameApi.ReviewUserName)
		gameRouter.POST("refuse", gameApi.RefuseUserName)

		gameRouter.POST("upload", questApi.Upload)
		gameRouter.GET("showGroup", questApi.ShowGroup)
		gameRouter.GET("questions", questApi.ShowQuestions)
		gameRouter.POST("setGameid", questApi.SetGameID)
		gameRouter.POST("createGroup", questApi.CreateGroup)
		gameRouter.GET("oneGroup", questApi.QuestionsById)
		gameRouter.POST("addsubject", questApi.AddQuestions)
		gameRouter.POST("deleteGroupSubject", questApi.DeleteQuestionInGroup)
		gameRouter.POST("updateQuestion", questApi.UpdateQuestion)
		gameRouter.POST("deleteQuestion", questApi.DeleteQuestion)

	}
	return R
}
