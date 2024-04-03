package cmd

import (
	"fmt"
	"os"

	"github.com/zc17375/e-portfolio-server/cmd/internal"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 取得 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.EP_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.EP_CONFIG.Zap.Director)
		_ = os.Mkdir(global.EP_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.EP_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
