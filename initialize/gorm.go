package initialize

import (
	"gameday/db/model"
	"gameday/global"
	"go.uber.org/zap"
	"os"
)

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.GameDB
	err := db.AutoMigrate(
		// 管理员表
		model.Admin{},
	)
	if err != nil {
		global.GameLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GameLog.Info("register table success")
}
