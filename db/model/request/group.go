package request

import "gameday/db/model"

// GroupRequest 用来接收设置比赛的题目组的数据
type GroupRequest struct {
	GroupId int `json:"group_id"`
	GameId  int `json:"game_id"`
}

type AllGame struct {
	Game      []model.Game
	GroupName string `json:"group_name"`
}
