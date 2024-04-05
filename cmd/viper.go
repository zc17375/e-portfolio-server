package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zc17375/e-portfolio-server/cmd/internal"
	"github.com/zc17375/e-portfolio-server/global"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				config = internal.ConfigDefaultFile
				fmt.Printf("您正在使用gin模式的%s環境,config的路徑為%s\n", gin.Mode(), internal.ConfigDefaultFile)
			case gin.ReleaseMode:
				config = internal.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s環境,config的路徑為%s\n", gin.Mode(), internal.ConfigReleaseFile)
			case gin.TestMode:
				config = internal.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s環境,config的路徑為%s\n", gin.Mode(), internal.ConfigTestFile)
			}
		} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
			config = configEnv
			fmt.Printf("您正在使用%s環境,config的路徑為%s\n", internal.ConfigEnv, config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()傳遞的值,config的路徑為%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.EP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.EP_CONFIG); err != nil {
		panic(err)
	}

	return v
}
