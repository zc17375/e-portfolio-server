package service

type ServiceGroup struct {
	AuthService
	HelloService
}

var ServiceGroupApp = new(ServiceGroup)
