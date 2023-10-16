package main

import (
	"gameday/core"
	"gameday/global"
	"gameday/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化viper(配置文件)
	global.GameVip = core.Viper()

	// 其他初始化

	// 初始化zap(日志系统)
	global.GameLog = core.Zap()
	zap.ReplaceGlobals(global.GameLog)
	// 初始化gorm(数据库-Mysql)
	global.GameDB = initialize.Mysql()
	if global.GameDB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭程序
		db, _ := global.GameDB.DB()
		defer db.Close()
	}
	//gin.DefaultWriter = io.Discard
	// 启动gin
	core.Run()

	// 编写测试
	//var user []model.User
	//var game *model.Game
	//global.GameDB.Preload("User").Where("id", 1).Find(&game)
	////_ = global.GameDB.Select("hashcode").Where("game_id = ?", 43).Find(&user)
	//fmt.Println(game)

	//fmt.Println(*user)
}
