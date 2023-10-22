package router

import (
	"gameday/router/admin"
	"gameday/router/game"
	"gameday/router/user"
)

type RouterGroup struct {
	Admin admin.RouterGroup
	Game  game.RouterGroup
	User  user.UserGroup
}

var RouterGroupApp = new(RouterGroup)
