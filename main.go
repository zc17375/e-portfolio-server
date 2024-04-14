package main

import (
	"github.com/zc17375/e-portfolio-server/cmd"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/initialize"
	"go.uber.org/zap"
)

// @title                       E-Portfolio Swagger API接口文件
// @version                     v1.0.0
// @description                 使用golang gin開發全端個人作品集網站
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化相關設定
	global.EP_VP = cmd.Viper() // 初始化Viper

	// 日誌配置初始化
	global.EP_LOG = cmd.Zap()
	zap.ReplaceGlobals(global.EP_LOG)
	global.EP_LOG.Info("Server相關設置初始化成功!!")

	// mysql數據庫配置初始化
	global.EP_DB = initialize.Gorm() // gorm連線資料庫
	if global.EP_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序結束前關閉數據庫連接
		db, _ := global.EP_DB.DB()
		defer db.Close()
	}

	// mongodb數據庫配置初始化
	global.EP_MongoDB = initialize.InitMongoDB()
	// mdb := initialize.InitMongoDB()

	// collection := mdb.Collection("students")
	// // 查询示例文档
	// var result map[string]interface{}
	// err := collection.FindOne(context.TODO(), map[string]interface{}{
	// 	"name": "Patrick",
	// }).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Found document:", result)
	defer initialize.Disconnect()

	// 執行伺服器
	cmd.RunServer()
	global.EP_LOG.Info("Server成功啟動..正在運行中!")

}
