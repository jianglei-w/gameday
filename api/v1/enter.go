package v1

import (
	"gameday/api/v1/admin"
	"gameday/api/v1/game"
	"gameday/api/v1/question"
	"gameday/api/v1/user"
)

type ApiGroup struct {
	AdminApiGroup admin.ApiGroup
	GameApiGroup  game.ApiGroup
	QuestApiGroup question.ApiGroup
	UserApiGroup  user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
