package router

// 註冊Router
type RouterGroup struct {
	AuthRouter
	IndividualRouter
	PortfolioRouter
}

var RouterGroupApp = new(RouterGroup)

