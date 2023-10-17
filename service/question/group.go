package question

import (
	"errors"
	"gameday/db/model"
	"gameday/db/model/request"
	"gameday/global"
)

func (qs *QuestionService) ShowGroup() (*[]model.Group, error) {
	var groups []model.Group

	if err := global.GameDB.Find(&groups).Error; err != nil {
		return nil, err
	}

	return &groups, nil
}

func (qs *QuestionService) CreateGroup(group *model.Group) (*model.Group, error) {
	if global.GameDB.Where("name = ?", group.Name).Find(&model.Group{}).RowsAffected == 0 {
		if err := global.GameDB.Create(&group).Error; err != nil {
			return nil, err
		}
		return group, nil
	}
	return nil, errors.New(group.Name + "组已存在")
}

func (qs *QuestionService) SetGameId(setGame *request.GroupRequest) error {
	// TODO暂不做判断

	err := global.GameDB.Model(&model.Game{}).Where("id", setGame.GameId).Update("group_id", setGame.GroupId).Error
	if err != nil {
		return err
	}

	return nil
}
