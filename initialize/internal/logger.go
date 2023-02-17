package internal

import (
	"fmt"
	"server/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer // struct内嵌套一个interface，表示在struct实例化时，需要一个实现此接口的struct进行赋值
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...any) {
	var logZap bool
	switch global.GGG_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GGG_CONFIG.Mysql.LogZap
	}
	if logZap {
		// 是否通过zap写入日志
		global.GGG_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}

}
