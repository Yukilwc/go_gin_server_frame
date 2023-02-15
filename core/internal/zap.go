package internal

import (
	"fmt"
	"server/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap *_zap = new(_zap)

type _zap struct {
}

func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotateLogs.GetWriteSyncer(l.String()) // 日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GGG_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {

	}
}
