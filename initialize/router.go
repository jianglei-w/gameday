package initialize

import (
	v1 "gameday/api/v1"
	"gameday/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	Router.POST("/admin/login", v1.ApiGroupApp.UserApiGroup.Login)

	return Router
}
