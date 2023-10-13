package v1

import "gameday/api/v1/user"

type ApiGroup struct {
	UserApiGroup user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
