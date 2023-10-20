package request

// GroupRequest 用来接收设置比赛的题目组的数据
type GroupRequest struct {
	GroupId int `json:"group_id"`
	GameId  int `json:"game_id"`
}

// QuestionGroup 用来操作many2many表
type QuestionGroup struct {
	QuestionID []uint `json:"subject_list"`
	GroupID    uint   `json:"group_id"`
}
