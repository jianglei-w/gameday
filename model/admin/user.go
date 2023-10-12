package admin

import "gameday/global"

type User struct {
	global.GameModel
	Username string `json:"userName" gorm:"index; comment:用户名`
	Password string `json:"passWord" gorm:"index; comment:密码`
}

func (u *User) TableName() string {
	return "admin"
}
