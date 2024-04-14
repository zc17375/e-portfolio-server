package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zc17375/e-portfolio-server/docs"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/middleware"
	"github.com/zc17375/e-portfolio-server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	routers := router.RouterGroupApp

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.EP_LOG.Info("register swagger handler")

	// 中間件跨域
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.EP_LOG.Info("use middleware cors")

	PublicGroup := Router.Group(global.EP_CONFIG.System.RouterPrefix)
	{
		// 健康檢測
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	PrivateGroup := Router.Group(global.EP_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth())
	{
		routers.GetUserRouter(PrivateGroup) // 注册用户路由
		routers.GetAuthRouter(PrivateGroup, PublicGroup)  // 註冊與登入路由，不檢查權限
		routers.GetIndividualRouter(PrivateGroup)
	}

	global.EP_LOG.Info("router register success")
	return Router
}
