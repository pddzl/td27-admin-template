package main

import (
	"go.uber.org/zap"

	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.TD27_VP = core.Viper() // 初始化viper
	global.TD27_LOG = core.Zap()  // 初始化zap日志
	zap.ReplaceGlobals(global.TD27_LOG)
	global.TD27_DB = initialize.Gorm() // gorm连接数据库
	if global.TD27_DB != nil {
		initialize.RegisterTables(global.TD27_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.TD27_DB.DB()
		defer db.Close()
	}
}
