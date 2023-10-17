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
	Name    string `json:"game_name"`   // 比赛名称
	People  uint   `json:"people"`      // 参与人数
	Status  bool   `json:"game_status"` // 0是初始化状态，1是开始比赛，2是暂停比赛，3是删除比赛
	GroupID uint   `json:"group_id"`
	Group   Group  `json:"group"` // 题目组ID
	Hashes  []User
}

func (g *Game) TableName() string {
	return "game"
}

type Group struct {
	gorm.Model
	Name     string     `json:"group_name"` // 题目组名称
	Question []Question // 包含哪些题目
	Games    []Game     `gorm:"foreignKey:GroupID"` // 属于哪些比赛
}

func (g *Group) TableName() string {
	return "group"
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

type Question struct {
	gorm.Model
	Title   string `json:"title"`    // 题目名 eg: 消失的Log
	Text    string `json:"text"`     // 题目介绍， eg: 找出xxx
	Answer  string `json:"answer"`   // 题目答案，eg: 可以是一串数字，字符等
	Grade   int    `json:"grade"`    // 分数：eg: 暂时未设计有哪些分数
	Level   string `json:"level"`    // 级别: eg: 暂定简单，普通，中等，困难
	GroupId uint   `json:"group_id"` // 组，eg: 属于什么组
}

func (q *Question) TableName() string {
	return "question"
}
