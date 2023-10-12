package global

import (
	"gameday/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GameConfig *config.Server
	GameVip    *viper.Viper
	GameLog    *zap.Logger
	GameDB     *gorm.DB
)
