package initialize

import (
	"server/config"
	"server/global"

	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GGG_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	if sysDB, ok := dbMap[sys]; ok {
		global.GGG_DB = sysDB
	}
	global.GGG_DBList = dbMap
}
