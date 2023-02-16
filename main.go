package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.GGG_VP = core.Viper()
	initialize.OtherInit()
	global.GGG_LOG = core.Zap()
	global.GGG_LOG.Info("Zap初始化完成")
	// global.GGG_LOG.DPanic("退出了")
}
