package config

// 服务端配置结构体 用来承载从yaml文件读入的配置
type Server struct {
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
}
