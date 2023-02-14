package initialize

import (
	"fmt"
	"server/global"
	"server/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GGG_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GGG_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	fmt.Println("BlackCache ", global.BlackCache)

}
