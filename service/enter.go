package service

type ServiceGroup struct {
	AuthService
	HelloService
	IndividualService
}

var ServiceGroupApp = new(ServiceGroup)
