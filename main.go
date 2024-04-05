package main

import (
	"github.com/zc17375/e-portfolio-server/cmd"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化相關設定
	global.EP_VP = cmd.Viper() // 初始化Viper

	// 日誌配置初始化
	global.EP_LOG = cmd.Zap()
	zap.ReplaceGlobals(global.EP_LOG)
	global.EP_LOG.Info("Server相關初始化成功!!")

	// 數據庫配置初始化
	global.EP_DB = initialize.Gorm() // gorm連線資料庫
	if global.EP_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.EP_DB.DB()
		defer db.Close()
	}

	// 執行伺服器
	cmd.RunServer()
	global.EP_LOG.Info("Server成功啟動..正在運行中!")

}
