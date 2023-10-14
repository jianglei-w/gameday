package router

import "gameday/router/user"

type RouterGroup struct {
	User user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
