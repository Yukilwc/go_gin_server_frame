package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
}

func (z *Zap) ZapEncoderLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
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
