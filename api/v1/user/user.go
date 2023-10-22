package user

import (
	"gameday/db/model"
	"gameday/db/model/response"
	"github.com/gin-gonic/gin"
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
