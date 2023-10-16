package game

import (
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

	//for i := 0; i < len(userList); i++ {
	//	userList[i].Game.Name = game.Name
	//}

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithData(userList, "创建成功", c)
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
