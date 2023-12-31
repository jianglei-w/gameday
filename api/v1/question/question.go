package question

import (
	"gameday/db/model"
	"gameday/db/model/request"
	"gameday/db/model/response"
	"gameday/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type QuestApi struct {
}

var questService = service.ServiceGroupApp.QuestService

func (q *QuestApi) Upload(c *gin.Context) {

	var quest model.Question

	c.ShouldBindJSON(&quest)

	if err := questService.Upload(&quest); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("上传成功", c)
}

func (q *QuestApi) ShowQuestions(c *gin.Context) {

	questions, err := questService.ShowQuestions()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithData(questions, "查询成功", c)
}

func (q *QuestApi) QuestionsById(c *gin.Context) {
	groupId, err := strconv.ParseUint(c.Query("group_id"), 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	group, err := questService.QuestionsById(uint(groupId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(group, "查询成功", c)

}

func (q *QuestApi) AddQuestions(c *gin.Context) {
	var qGroup *request.QuestionGroup
	err := c.ShouldBindJSON(&qGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = questService.AddQuestions(qGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("添加成功", c)
}

func (q *QuestApi) UpdateQuestion(c *gin.Context) {
	var question *model.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = questService.UpdateQuestion(question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("更新成功", c)

}

func (q *QuestApi) DeleteQuestion(c *gin.Context) {
	var question *model.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = questService.DeleteQuestion(question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("删除成功", c)

}
