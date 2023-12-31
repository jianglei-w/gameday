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
	Status  uint   `json:"game_status"` // 0是初始化状态，1是开始比赛，2是暂停比赛，3是完成比赛
	GroupID uint   `json:"group_id"`
	Group   Group  `json:"group"` // 题目组ID
	Hashes  []User
}

func (g *Game) TableName() string {
	return "game"
}

type Group struct {
	gorm.Model
	Name      string     `json:"group_name"`                 // 题目组名称
	Questions []Question `gorm:"many2many:question_groups;"` // 包含哪些题目
	Games     []Game     `gorm:"foreignKey:GroupID"`         // 属于哪些比赛
}

func (g *Group) TableName() string {
	return "group"
}

type User struct {
	gorm.Model
	//score = db.relationship('Score', backref='admin')
	//event = db.relationship('Event', backref='admin')
	Hashcode string `json:"hashcode"`
	Username string `json:"username"`
	Status   uint   `json:"status" ` // 名字审核状态0：未设置，1: 待审核，2：审核通过，3: 已驳回
	GameID   uint   `json:"game_id"`
	Score    uint   `json:"score"`
}

func (u *User) TableName() string {
	return "user"
}

type Question struct {
	gorm.Model
	Title   string  `json:"title"`    // 题目名 eg: 消失的Log
	Text    string  `json:"text"`     // 题目介绍， eg: 找出xxx
	Answer  string  `json:"answer"`   // 题目答案，eg: 可以是一串数字，字符等
	Grade   int     `json:"grade"`    // 分数：eg: 暂时未设计有哪些分数
	Level   string  `json:"level"`    // 级别: eg: 暂定简单，普通，中等，困难
	GroupID uint    `json:"group_id"` // 组，eg: 属于什么组
	Group   []Group `gorm:"many2many:question_groups;"`
}

func (q *Question) TableName() string {
	return "question"
}

// Complete 记录哪些用户完成了哪些题目
type Complete struct {
	gorm.Model
	UserId     uint `json:"user_id"`     // 用户id
	QuestionId uint `json:"question_id"` // 题目id
	Score      uint `json:"score"`       // 题目分数
}

func (s *Complete) TableName() string {
	return "complete"
}

// Score 记录用户得分情况
