package response

import "gameday/db/model"

type GameResponse struct {
	GameName string `json:"game_name"`
	//People   int                      `json:"people"`
	HashList []*model.User `json:"hash_list"`
}
