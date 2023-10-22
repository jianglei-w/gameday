package admin

import "gameday/router/user"

type RouterGroup struct {
	user.UserRouter
	AdminRouter
	TestRouter
}
