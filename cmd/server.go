package cmd

import (
	"fmt"
	"time"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// // 从db加载jwt数据
	// if global.EP_DB != nil {
	// 	system.LoadAll()
	// }

	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")

	port := fmt.Sprintf(":%d", global.EP_CONFIG.System.Port)
	s := initServer(port, Router)
	time.Sleep(10 * time.Microsecond)
	global.EP_LOG.Info("server run success on ", zap.String("address", port))

	fmt.Printf(`
	%s e-portfolio-server
	%s: v1.0.0
	自動化文件網址:http://127.0.0.1%s/swagger/index.html
	前端網址:http://127.0.0.1:8080`,
		"Project",
		"Current Version", port)
	global.EP_LOG.Error(s.ListenAndServe().Error())
}
