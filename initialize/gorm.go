package initialize

import (
	"os"
	"server/global"
	"server/system"

	"go.uber.org/zap"
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

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
	)
	if err != nil {
		global.GGG_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GGG_LOG.Info("register table success")
}
