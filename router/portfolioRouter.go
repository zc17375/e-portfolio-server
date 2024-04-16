package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zc17375/e-portfolio-server/api/v1/controller"
)

type PortfolioRouter struct{}

func (s *AuthRouter) GetPortfolioRouter(Router *gin.RouterGroup) {
	// 取使用者相關紀錄
	portfolioRouter := Router.Group("portfolio")

	portfolioApi := v1.ApiGroupApp.PortfolioApi
	{
		portfolioRouter.GET(":userName", portfolioApi.GetUserPortfolio) //新增個人作品相關資料
		portfolioRouter.POST("list", portfolioApi.GetAllPortfolio) //更新個人作品相關資料
	}
}

