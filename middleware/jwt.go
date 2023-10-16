package middleware

import (
	"errors"
	"gameday/db/model/response"
	"gameday/utils"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithMessage("授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
