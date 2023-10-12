package initialize

import (
	"gameday/global"
	"gameday/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	adminRouter := router.RouterGroupApp.Admin
	//userRouter := router.RouterGroupApp.User

	PublicGroup := Router.Group(global.GameConfig.System.RouterPrefix)
	{
		PublicGroup.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
	}
	{

	}
	PrivateGroup := Router.Group(global.GameConfig.System.RouterPrefix)
	{
		adminRouter.Admin.InitAdminRouter(PrivateGroup)
	}
	return Router
}
