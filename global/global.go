package global

import (
	"server/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GGG_VP     *viper.Viper // 全局Viper
	GGG_LOG    *zap.Logger
	GGG_CONFIG config.Server // 存放从yaml配置文件读取的结构体
)
