package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zc17375/e-portfolio-server/docs"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	routers := router.RouterGroupApp

	Router.GET(global.EP_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.EP_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.EP_CONFIG.System.RouterPrefix)
	{
		// 健康檢測
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	routers.GetAuthRouter(PublicGroup) // 註冊與登入路由，不檢查權限
	// PrivateGroup := Router.Group(global.EP_CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
	// systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
	// routers.GetUserRouter(PrivateGroup)                   // 注册用户路由
	
	}

	global.EP_LOG.Info("router register success")
	return Router
}

