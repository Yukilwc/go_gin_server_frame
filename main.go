package main

import (
	"server/core"
	"server/global"
	"server/initialize"

	"go.uber.org/zap"
)

func main() {
	global.GGG_VP = core.Viper()
	initialize.OtherInit()
	global.GGG_LOG = core.Zap()
	zap.ReplaceGlobals(global.GGG_LOG)
}
