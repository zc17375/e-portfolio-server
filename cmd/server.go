package cmd

import (
	"fmt"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/initialize"
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
	fmt.Println(s.ListenAndServe().Error())
	// global.EP_LOG.Error(s.ListenAndServe().Error())
}
