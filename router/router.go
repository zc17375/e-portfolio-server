package router

// 註冊Router
type RouterGroup struct {
	AuthRouter
	// HelloRouter
	// UserRouter
}

var RouterGroupApp = new(RouterGroup)

