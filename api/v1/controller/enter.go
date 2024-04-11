package controller

import (
	"github.com/zc17375/e-portfolio-server/service"
)

type ApiGroup struct {
	AuthApi
	HelloApi
	UserApi
}

var (
	authService = service.ServiceGroupApp.AuthService
	helloService = service.ServiceGroupApp.HelloService
	// userService = service.ServiceGroupApp.UserService
)

var ApiGroupApp = new(ApiGroup)