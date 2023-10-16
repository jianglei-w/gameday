package question

import (
	"gameday/db/model"
	"gameday/db/model/response"
	"gameday/service"
	"github.com/gin-gonic/gin"
)

type QuestApi struct {
}

func (q *QuestApi) Upload(c *gin.Context) {
	questService := service.ServiceGroupApp.QuestService
	var quest model.Question

	c.ShouldBindJSON(&quest)

	if err := questService.Upload(&quest); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("上传成功", c)
}
