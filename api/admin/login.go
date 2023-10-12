package admin

import "github.com/gin-gonic/gin"

type User struct {
}

func (u *User) Login(c *gin.Context) {
	c.JSON(200, "I am Login")
}
