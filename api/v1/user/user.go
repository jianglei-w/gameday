package user

import (
	"gameday/db/model"
	"gameday/db/model/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (u *UserApi) UpdateUsername(c *gin.Context) {
	var r model.User
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.UpdateUsername(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功，等待审核", c)
}

func (u *UserApi) GetUserByHash(c *gin.Context) {
	var r model.User
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.GetUserByHash(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithData(user, "返回成功", c)

}

func (u *UserApi) GetSuccessQuestionByUserId(c *gin.Context) {
	var r model.Complete
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.GetSuccessQuestionByUserId(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithData(user, "查询成功", c)
}

func (u *UserApi) UserScore(c *gin.Context) {
	var r model.Complete
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sumScore, err := userService.UserScore(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithData(sumScore, "查询成功", c)
}

func (u *UserApi) StopGame(c *gin.Context) {
	a := struct {
		Gameid uint `json:"game_id"`
	}{}
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	game, err := userService.StopGame(a.Gameid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(game, "查询成功", c)
}

func (u *UserApi) PostAnswer(c *gin.Context) {
	var quest model.Question
	err := c.ShouldBindBodyWith(&quest, binding.JSON)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.PostAnswer(&quest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userid := struct {
		UserId uint `json:"user_id"`
	}{}
	err = c.ShouldBindBodyWith(&userid, binding.JSON)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.StoreComplete(userid.UserId, quest.ID, uint(quest.Grade))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("提交成功", c)
}

func (u *UserApi) RankingList(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := userService.RankingList(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(result, "查询成功", c)
}

func (u *UserApi) UserThing(c *gin.Context) {
	var user model.Complete
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	results, err := userService.UserThing(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(results, "查询成功", c)
}

func (u *UserApi) Event(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	results, err := userService.Event(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(results, "查询成功", c)
}
