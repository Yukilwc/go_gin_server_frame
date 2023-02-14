package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.GGG_VP = core.Viper()
	initialize.OtherInit()
}
