package user

import (
	"gameday/db/model"
	myjwt "gameday/db/model/jwt"
	"gameday/db/model/request"
	"gameday/db/model/response"
	"gameday/global"
	"gameday/service"
	"gameday/utils"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

// Login
//
//	@Tags		User
//	@Summary	用户登录
//	@Produce	json
//	@Param		data	body	request.Login	true "用户名, 密码"
//	@Router		/admin/login [post]
func (b *BaseApi) Login(c *gin.Context) {

	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &model.Admin{Username: l.Username, Password: l.Password}
	user, err := service.ServiceGroupApp.UserService.Login(u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
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
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(response.LoginResponse{
		Id:        user.ID,
		UserName:  user.Username,
		Password:  user.Password,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}
