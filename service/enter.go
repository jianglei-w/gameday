package service

import (
	"gameday/service/game"
	"gameday/service/question"
	"gameday/service/user"
)

type serviceGroup struct {
	UserService  user.ServiceGroup
	GameService  game.ServiceGroup
	QuestService question.ServiceGroup
}

var ServiceGroupApp = new(serviceGroup)
