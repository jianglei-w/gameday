package core

import (
	"fmt"
	"gameday/core/internal"
	"gameday/global"
	"gameday/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {

	if ok, _ := utils.PathExists(global.GameConfig.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GameConfig.Zap.Director)
		_ = os.Mkdir(global.GameConfig.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GameConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}