package service

type ServiceGroup struct {
	AuthService
	HelloService
	IndividualService
	PortfolioService
}

var ServiceGroupApp = new(ServiceGroup)
