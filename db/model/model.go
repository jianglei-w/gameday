package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Admin) TableName() string {
	return "admin"
}

type Game struct {
	gorm.Model
	Name   string `json:"game_name"`   // 比赛名称
	People uint   `json:"people"`      // 参与人数
	Status bool   `json:"game_status"` // 0是初始化状态，1是开始比赛，2是暂停比赛，3是删除比赛
	Hashes []User
}

func (g *Game) TableName() string {
	return "game"
}

type User struct {
	gorm.Model
	//score = db.relationship('Score', backref='user')
	//event = db.relationship('Event', backref='user')
	Hashcode string `json:"hashcode"`
	Username string `json:"username"`
	Status   bool   `json:"status" ` // 名字审核状态
	GameID   uint   `json:"game_id"`
}

func (u *User) TableName() string {
	return "user"
}

//type Question struct {
//	gorm.Model
//
//}
