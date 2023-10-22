package user

import (
	"gameday/db/model"
	"gameday/global"
)

func (us *UserService) UpdateUsername(user *model.User) error {
	err := global.GameDB.Model(&model.User{}).Where("hashcode = ?", user.Hashcode).Updates(map[string]interface{}{
		"username": user.Username,
		"status":   1,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByHash(user *model.User) (*model.User, error) {
	if err := global.GameDB.Model(&model.User{}).Where("hashcode", user.Hashcode).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetSuccessQuestionByUserId(s *model.Complete) ([]uint, error) {
	//var score []*model.Score
	var completeIDs []uint
	if err := global.GameDB.
		Model(&model.Complete{}).
		Where("user_id = ?", s.UserId).
		Pluck("question_id", &completeIDs).Error; err != nil {
		return nil, err
	}

	return completeIDs, nil
}

func (us *UserService) UserScore(s *model.Complete) (uint, error) {
	var score uint
	if err := global.GameDB.Model(&model.Complete{}).
		Select("COALESCE(SUM(score), 0)"). // 默认值，如果字段sum(score)为NULL则设置默认值为0
		Where("user_id = ?", s.UserId).
		Scan(&score).
		Error; err != nil {
		return 0, err
	}

	return score, nil
}
