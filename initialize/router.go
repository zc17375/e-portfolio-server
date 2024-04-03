package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	routers := router.RouterGroupApp

	PublicGroup := Router.Group("")
	{
		// 健康檢測
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		routers.InitHelloRouter(PublicGroup)
	}

	return Router
}
