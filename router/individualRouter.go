package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zc17375/e-portfolio-server/api/v1/controller"
)

type IndividualRouter struct{}

func (s *AuthRouter) GetIndividualRouter(Router *gin.RouterGroup) {
	// 取使用者相關紀錄
	individualRouter := Router.Group("individual")

	individualApi := v1.ApiGroupApp.IndividualApi
	{
		individualRouter.POST("create", individualApi.CreateIndividual)               // 會員登入
	}
}
