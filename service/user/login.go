package user

import (
	"errors"
	"gameday/db/model"
	"gameday/global"
)

type UserService struct {
}

func (us *UserService) Login(user *model.User) (*model.User, error) {
	var u model.User
	if result := global.GameDB.Where("hashcode = ?", user.Hashcode).First(&u); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("没有该用户")
	}

	return &u, nil
}
