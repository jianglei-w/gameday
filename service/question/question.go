package question

import (
	"errors"
	"gameday/db/model"
	"gameday/db/model/request"
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

func (qs *QuestionService) AddQuestions(qGroup *request.QuestionGroup) error {
	type QuestionGroup struct {
		GroupID    uint
		QuestionID uint
	}

	for _, questionId := range qGroup.QuestionID {
		err := global.GameDB.Create(&QuestionGroup{
			GroupID:    qGroup.GroupID,
			QuestionID: questionId,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (qs *QuestionService) DeleteQuestionInGroup(groupId, questionId uint) error {
	type QuestionGroup struct {
		GroupID    uint
		QuestionID uint
	}

	if err := global.GameDB.Where(&QuestionGroup{GroupID: groupId, QuestionID: questionId}).Delete(&QuestionGroup{GroupID: groupId, QuestionID: questionId}).Error; err != nil {
		return err
	}

	return nil
}

func (qs *QuestionService) UpdateQuestion(question *model.Question) error {
	if err := global.GameDB.Model(&model.Question{}).Where("id = ?", question.ID).Updates(&question).Error; err != nil {
		return err
	}

	return nil
}

func (qs *QuestionService) DeleteQuestion(question *model.Question) error {
	if err := global.GameDB.Delete(&question).Error; err != nil {
		return err
	}

	return nil
}
