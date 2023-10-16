package router

import (
	"gameday/router/game"
	"gameday/router/user"
)

type RouterGroup struct {
	User user.RouterGroup
	Game game.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
