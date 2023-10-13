package user

import (
	"errors"
	"fmt"
	"gameday/db/model"
	"gameday/global"
)

type UserService struct {
}

func (u *UserService) Login(admin *model.Admin) (userInter *model.Admin, err error) {
	if global.GameDB == nil {
		return nil, fmt.Errorf("do not init")
	}
	var user model.Admin
	if err = global.GameDB.Where("username = ?", admin.Username).First(&user).Error; err == nil {
		// TODO 使用对比方法对比密码
		if admin.Password == user.Password {
			return &user, nil
		} else {
			return nil, errors.New("密码错误")
		}
	}
	return nil, err
}