package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zc17375/e-portfolio-server/api/v1/controller"
)

type AuthRouter struct{}

func (s *AuthRouter) GetAuthRouter(Router *gin.RouterGroup) {
	// 取使用者相關紀錄
	userRouter := Router.Group("auth")
	//.Use(middleware.OperationRecord())
	// 不取使用者相關紀錄
	// userRouterWithoutRecord := Router.Group("user")

	authApi := v1.ApiGroupApp.AuthApi
	{
		userRouter.POST("login", authApi.Login)               // 會員登入
	}
}
