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
		individualRouter.POST("create", individualApi.CreateIndividual)   //新增個人作品相關資料
		individualRouter.POST("update", individualApi.UpdateIndividual)   //更新個人作品相關資料
		individualRouter.DELETE("delete", individualApi.DeleteIndividual) //刪除個人作品相關資料
	}
}
