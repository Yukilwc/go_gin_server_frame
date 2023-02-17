package config

// 服务端配置结构体 用来承载从yaml文件读入的配置
type Server struct {
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
