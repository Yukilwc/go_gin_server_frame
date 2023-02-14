package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"server/core/internal"
	"server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	// 选择配置文件
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
					// 默认Mode就是debug
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			}
		}

	}
	v := initViper(config)
	return v
}

func initViper(config string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 读入
	if err = v.Unmarshal(&global.GGG_CONFIG); err != nil {
		fmt.Println(err)
	}
	// 监听文件变动
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 变动时，重新读入配置
		if err = v.Unmarshal(&global.GGG_CONFIG); err != nil {
			fmt.Println(err)
		}

	})
	// 自动识别加载根路径
	global.GGG_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
