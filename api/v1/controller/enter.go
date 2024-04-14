package controller

import (
	"github.com/zc17375/e-portfolio-server/service"
)

type ApiGroup struct {
	AuthApi
	HelloApi
	UserApi
	IndividualApi
}

var (
	authService = service.ServiceGroupApp.AuthService
	helloService = service.ServiceGroupApp.HelloService
	individualService = service.ServiceGroupApp.IndividualService
)

var ApiGroupApp = new(ApiGroup)