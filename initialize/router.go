package initialize

import (
	_ "gameday/docs"
	"gameday/global"
	"gameday/middleware"
	"gameday/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	UserRouter := router.RouterGroupApp.User

	// 注册swag
	Router.GET(global.GameConfig.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GameLog.Info("register swag Success")

	PublicGroup := Router.Group(global.GameConfig.System.RouterPrefix) // 公共路由的/health路由,健康监测
	{
		PublicGroup.GET("health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	PrivateGroup := Router.Group(global.GameConfig.System.RouterPrefix)
	PrivateGroup.Use(middleware.Cors()) // 中间件
	{
		UserRouter.InitUserRouter(PrivateGroup) // user类别路由
	}

	global.GameLog.Info("router register success")

	return Router
}
