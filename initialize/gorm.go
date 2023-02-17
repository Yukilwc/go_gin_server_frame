package initialize

import (
	"server/global"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GGG_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
