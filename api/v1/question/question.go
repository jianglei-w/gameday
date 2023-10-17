package question

import (
	"gameday/db/model"
	"gameday/db/model/response"
	"gameday/service"
	"github.com/gin-gonic/gin"
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
	}

	response.OKWithData(questions, "查询成功", c)
}
