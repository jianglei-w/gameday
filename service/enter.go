package service

import (
	"gameday/service/game"
	"gameday/service/user"
)

type serviceGroup struct {
	UserService user.ServiceGroup
	GameService game.ServiceGroup
}

var ServiceGroupApp = new(serviceGroup)
