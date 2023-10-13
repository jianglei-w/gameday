package user

import "gameday/service"

type ApiGroup struct {
	BaseApi
}

var userService = service.ServiceGroupApp.UserService
