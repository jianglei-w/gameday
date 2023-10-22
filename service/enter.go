package service

import (
	"gameday/service/admin"
	"gameday/service/game"
	"gameday/service/question"
	"gameday/service/user"
)

type serviceGroup struct {
	AdminService admin.ServiceGroup
	GameService  game.ServiceGroup
	QuestService question.ServiceGroup
	UserService  user.ServiceGroup
}

var ServiceGroupApp = new(serviceGroup)
