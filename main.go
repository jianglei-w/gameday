package main

import (
	"gameday/core"
	"gameday/global"
	"gameday/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
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
	gin.DefaultWriter = io.Discard
	// 启动gin
	core.Run()

	// 编写测试
}
