package controller

import (
	"github.com/zc17375/e-portfolio-server/service"
)

type ApiGroup struct {
	HelloApi
}

var (
	helloService = service.ServiceGroupApp.HelloService
)

var ApiGroupApp = new(ApiGroup)