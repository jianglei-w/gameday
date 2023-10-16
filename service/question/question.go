package question

import (
	"errors"
	"gameday/db/model"
	"gameday/global"
)

type QuestionService struct {
}

func (qs *QuestionService) Upload(quest *model.Question) error {
	if global.GameDB.Where("title = ?", quest.Title).First(&model.Question{}).RowsAffected == 0 {
		if err := global.GameDB.Create(&quest).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("题目" + quest.Title + "已存在")
}
