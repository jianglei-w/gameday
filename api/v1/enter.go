package v1

import (
	"gameday/api/v1/game"
	"gameday/api/v1/question"
	"gameday/api/v1/user"
)

type ApiGroup struct {
	UserApiGroup  user.ApiGroup
	GameApiGroup  game.ApiGroup
	QuestApiGroup question.QuestApi
}

var ApiGroupApp = new(ApiGroup)
