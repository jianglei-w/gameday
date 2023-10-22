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
	AdminRouter := router.RouterGroupApp.Admin
	GameRouter := router.RouterGroupApp.Game
	UserRouter := router.RouterGroupApp.User

	// 注册swag
	Router.GET(global.GameConfig.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GameLog.Info("register swag Success")

	// 全部路由跨域=_=
	Router.Use(middleware.Cors())

	PublicGroup := Router.Group(global.GameConfig.System.RouterPrefix) // 公共路由的/health路由,健康监测
	{
		PublicGroup.GET("health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")

		})

	}
	{ // 不需要鉴权的公共接口 login
		AdminRouter.InitAdminRouter(PublicGroup)
		UserRouter.InitUserRouter(PublicGroup)
	}

	// 需要鉴权的私有接口
	PrivateGroup := Router.Group(global.GameConfig.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWT()) // 中间件
	{
		// admin类别路由
		AdminRouter.InitTestRouter(PrivateGroup)

		// game类别路由
		GameRouter.InitGameRouters(PrivateGroup)

		// user类别路由
		UserRouter.InitAuthRouters(PrivateGroup)
	}

	global.GameLog.Info("router register success")

	return Router
}
