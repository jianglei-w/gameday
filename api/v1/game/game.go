package game

import (
	"fmt"
	"gameday/db/model"
	"gameday/db/model/request"
	"gameday/db/model/response"
	"gameday/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GameApi struct {
}

var gameService = service.ServiceGroupApp.GameService

// CreateGame
//
//	@Tags		Game
//	@Summary	创建游戏和生成hash
//	@Produce	json
//	@Param		data	body	request.Game	true "用户名, 密码"
//	@Router		/admin/createGame [post]
func (g *GameApi) CreateGame(c *gin.Context) {
	var r request.Game
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	game := &model.Game{
		Name:   r.GameName,
		People: r.People,
	}
	userList, err := gameService.CreateGame(game)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(userList, "创建成功", c)
}

func (g *GameApi) StartGame(c *gin.Context) {
	var r model.Game
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(r.ID)
	err = gameService.StartGame(r.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("比赛已经开始", c)
}

func (g *GameApi) StopGame(c *gin.Context) {
	var r model.Game
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gameService.StopGame(r.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("比赛已经暂停", c)
}

func (g *GameApi) DoneGame(c *gin.Context) {
	var r model.Game
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gameService.DoneGame(r.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("比赛已完成", c)
}
func (g *GameApi) LastGame(c *gin.Context) {

	users, err := gameService.LastGame()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(users, "操作成功", c)
}

func (g *GameApi) GetHash(c *gin.Context) {

	gameId, err := strconv.ParseUint(c.Query("game_id"), 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	users, err := gameService.GetHash(uint(gameId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(users, "查询成功", c)
}

func (g *GameApi) GetAllGame(c *gin.Context) {

	games, err := gameService.GetAllGame()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithData(games, "查询成功", c)
}

// RankList 排行榜
func (g *GameApi) RankList(c *gin.Context) {
	gameId, err := strconv.ParseUint(c.Query("game_id"), 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	game, err := gameService.RankList(uint(gameId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(game, "查询成功", c)
}

func (g *GameApi) ReviewUserName(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = gameService.ReviewUserName(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("用户名已通过", c)
}

func (g *GameApi) RefuseUserName(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = gameService.RefuseUserName(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OKWithMessage("用户已驳回", c)
}
