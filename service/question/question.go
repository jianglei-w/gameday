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

func (qs *QuestionService) ShowQuestions() ([]*model.Question, error) {
	var questions []*model.Question
	if err := global.GameDB.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (qs *QuestionService) QuestionsById(GroupID uint) (*model.Group, error) {
	var group *model.Group

	if err := global.GameDB.Preload("Questions").Where("id = ?", GroupID).Find(&group).Error; err != nil {
		return nil, err
	}

	return group, nil
}
