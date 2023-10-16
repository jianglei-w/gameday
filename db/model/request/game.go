package request

type Game struct {
	GameName string `json:"gamename"`
	People   uint   `json:"people"` // 参与人数
}
