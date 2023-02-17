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
	global.GGG_DB = initialize.Gorm()
	// initialize.Timer()
	initialize.DBList()
	if global.GGG_DB != nil {
		db, _ := global.GGG_DB.DB()
		defer db.Close()
	}
	// core.RunWindowsServer()
}
