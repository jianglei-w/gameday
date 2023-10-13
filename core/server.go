package core

import (
	"fmt"
	"gameday/global"
	"gameday/initialize"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run() {

	// 初始化路由
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GameConfig.System.Addr)
	s := initServer(address, Router)
	global.GameLog.Info("server run success on ", zap.String("address", address))

	if err := s.ListenAndServe(); err != nil && err.Error() != fmt.Sprintf("accept tcp [::]%s: use of closed network connection", address) {
		global.GameLog.Error("failed to listen gameday server: ", zap.Error(err))
	}
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
