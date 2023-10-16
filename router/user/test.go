package user

import (
	"fmt"
	myjwt "gameday/db/model/jwt"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func (tr *TestRouter) InitTestRouter(Router *gin.RouterGroup) (R *gin.IRoutes) {
	testRouter := Router.Group("test")
	{
		testRouter.GET("", func(c *gin.Context) {
			value, _ := c.Get("claims")
			claims := value.(*myjwt.CustomClaims)
			fmt.Println(claims)
			c.String(200, "test")
		})
	}
	return R
}
