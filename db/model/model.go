package model

import (
	"gameday/db/model/gameday"
)

type Admin struct {
	gameday.GameModel
	Username string `json:"username" gorm:"index"`
	Password string `json:"password" gorm:"index"`
}

func (a *Admin) TableName() string {
	return "admin"
}
