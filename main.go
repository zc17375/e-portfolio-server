package main

import (
	"github.com/zc17375/e-portfolio-server/cmd"
	"github.com/zc17375/e-portfolio-server/global"
	"go.uber.org/zap"
)

func main() {
	// 初始化相關設定

	// 日誌配置
	global.EP_LOG = cmd.Zap()
	zap.ReplaceGlobals(global.EP_LOG)

	global.EP_LOG.Info("router register success")

	// 執行伺服器
	cmd.RunServer()
}
