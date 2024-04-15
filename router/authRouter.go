package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zc17375/e-portfolio-server/api/v1/controller"
)

type AuthRouter struct{}

func (s *AuthRouter) GetAuthRouter(RouterPrt *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	// 取使用者相關紀錄
	NonAuthRouter := RouterPub.Group("auth")
	AuthRouter := RouterPrt.Group("auth")
	//.Use(middleware.OperationRecord())
	// 不取使用者相關紀錄
	// userRouterWithoutRecord := Router.Group("user")

	authApi := v1.ApiGroupApp.AuthApi
	{
		NonAuthRouter.POST("login", authApi.Login)            // 會員登入
		NonAuthRouter.POST("register", authApi.Register)      // 會員註冊
		AuthRouter.GET("refresh-token", authApi.RefreshToken) //刷新TOKEN
		AuthRouter.GET("delete", authApi.DeleteUser)          //註銷會員
	}
}
