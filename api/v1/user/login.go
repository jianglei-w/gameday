package user

import (
	"gameday/db/model"
	myjwt "gameday/db/model/jwt"
	"gameday/db/model/response"
	"gameday/global"
	"gameday/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) Login(c *gin.Context) {

	var l *model.Admin
	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"data": map[string]interface{}{},
			"msg":  err.Error(),
		})
		return
	}
	u := &model.Admin{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 7,
			"data": map[string]interface{}{},
			"msg":  "用户名不存在或者密码错误",
		})
		return
	}
	b.TokenNext(c, user)
	return

}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user *model.Admin) {
	j := &utils.JWT{
		SigningKey: []byte(global.GameConfig.JWT.SigningKey),
	}
	claims := j.CreateClaims(myjwt.BaseClaims{
		ID:       user.ID,
		Username: user.Username,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": map[string]interface{}{},
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response.LoginResponse{
			Id:        user.ID,
			UserName:  user.Username,
			Password:  user.Password,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		},
		"msg":    "Success",
		"status": 200,
	})
}
