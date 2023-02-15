package core

import (
	"fmt"
	"os"
	"server/global"
	"server/utils"

	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	if ok, _ := utils.PathExists(global.GGG_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GGG_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GGG_CONFIG.Zap.Director, os.ModePerm)
	}
}
