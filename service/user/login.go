package user

import (
	"errors"
	"fmt"
	"gameday/db/model"
	"gameday/global"
	"gorm.io/gorm"
)

type userService struct {
}

func (u *userService) Login(admin *model.Admin) (userInter *model.Admin, err error) {
	if global.GameDB == nil {
		return nil, fmt.Errorf("do not init")
	}
	var user model.Admin
	if err = global.GameDB.Where("username = ?", admin.Username).First(&user).Error; err == nil {
		// TODO 使用（加密）对比方法对比密码
		if admin.Password == user.Password {
			return &user, nil
		} else {
			return nil, errors.New("密码错误")
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	return nil, err
}
