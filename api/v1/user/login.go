package user

import (
	"gameday/db/model"
	myjwt "gameday/db/model/jwt"
	"gameday/db/model/response"
	"gameday/global"
	"gameday/service"
	"gameday/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.UserService
var gameService = service.ServiceGroupApp.GameService

func (u *UserApi) Login(c *gin.Context) {
	var r model.User
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, err := userService.Login(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 能查到数据，说明一定存在这个hashcode，则登录成功

	u.TokenNext(c, user)

}

func (u *UserApi) TokenNext(c *gin.Context, user *model.User) {
	j := &utils.JWT{
		SigningKey: []byte(global.GameConfig.JWT.SigningKey),
	}
	claims := j.CreateClaims(myjwt.BaseClaims{
		ID:       user.ID,
		Username: user.Hashcode,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	game, err := gameService.GetHash(user.GameID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msg := ""
	switch game.Status {
	case 0:
		msg = "比赛未开始"
	case 1:
		msg = "比赛已开始"
	case 2:
		msg = "比赛已暂停"
	case 3:
		msg = "比赛已结束"
	}
	response.OKWithData(response.UserLoginResponse{
		Id:        user.ID,
		HashCode:  user.Hashcode,
		Token:     token,
		GameId:    game.ID,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, msg, c)
}
