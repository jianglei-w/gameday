package question

import (
	"gameday/db/model"
	"gameday/db/model/request"
	"gameday/db/model/response"
	"gameday/service"
	"github.com/gin-gonic/gin"
)

var questApi = service.ServiceGroupApp.QuestService

func (q *QuestApi) ShowGroup(c *gin.Context) {

	group, err := questApi.ShowGroup()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OKWithData(group, "查询成功", c)

}

func (q *QuestApi) CreateGroup(c *gin.Context) {
	var group model.Group
	err := c.ShouldBindJSON(&group)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	createGroup, err := questApi.CreateGroup(&group)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OKWithData(createGroup, "查询成功", c)
}

func (q *QuestApi) SetGameID(c *gin.Context) {
	var SetGameId request.GroupRequest
	err := c.ShouldBindJSON(&SetGameId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = questApi.SetGameId(&SetGameId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)

	}
	response.OKWithMessage("设置成功", c)
}

func (q *QuestApi) DeleteQuestionInGroup(c *gin.Context) {
	type question struct {
		GroupID    uint `json:"group_id"`
		QuestionID uint `json:"question_id"`
	}
	var quest question
	err := c.ShouldBindJSON(&quest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = questApi.DeleteQuestionInGroup(quest.GroupID, quest.QuestionID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("移除成功", c)
}
