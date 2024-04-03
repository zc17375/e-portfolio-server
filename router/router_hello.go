package router

import (

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/api/v1/controller"
)


type HelloRouter struct{}

func (r *HelloRouter) InitHelloRouter(Router *gin.RouterGroup) {
	helloRouter := Router.Group("hello")
	helloApi := controller.ApiGroupApp.HelloApi
	{
		helloRouter.GET("/", helloApi.Hello)
	}
}
