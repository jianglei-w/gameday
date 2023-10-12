package api

import (
	"gameday/api/admin"
	"gameday/api/user"
)

type ApiGroup struct {
	Admin admin.ApiGroup
	User  user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
