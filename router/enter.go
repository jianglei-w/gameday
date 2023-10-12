package router

import (
	"gameday/router/admin"
	"gameday/router/user"
)

type RouterGroup struct {
	Admin admin.RouterGroup
	User  user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
