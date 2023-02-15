package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Director     string `mapstructure:"director" json:"director" yaml:"director"`
	MaxAge       int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
}

// 把Zap中的字符串level转换为zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel

	}

}
