package service

import "gameday/service/user"

type ServiceGroup struct {
	UserService user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
