package core

import (
	"fmt"
	"os"
	"server/core/internal"
	"server/global"
	"server/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GGG_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GGG_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GGG_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.GGG_CONFIG.Zap.ShowLine {
		logger.WithOptions(zap.AddCaller())
	}
	return logger
}
