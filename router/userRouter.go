package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zc17375/e-portfolio-server/api/v1/controller"
)

type UserRouter struct{}

func (s *AuthRouter) GetUserRouter(Router *gin.RouterGroup) {
	// 取使用者相關紀錄
	userRouter := Router.Group("user")

	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.GET("portfolio", userApi.Portfolio)               // 會員登入
	}
}
