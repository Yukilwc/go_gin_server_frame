package core

import (
	"flag"
	"fmt"
	"os"
	"server/core/internal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) != 0 {
		// 存在参数设定，优先级最高
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	} else {
		// 接收命令行指令的参数
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config != "" {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		} else {
			// 命令行没有配置，则走环境变量配置
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv != "" {
				// 走环境变量配置
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			} else {
				// 最后，从Gin的mode取
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDebugFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				}
			}
		}

	}
	return nil
}
