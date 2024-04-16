package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model/common"
	"go.uber.org/zap"
)

type PortfolioApi struct{}

// User Portfolio
// @Tags     Portfolio
// @Summary  個人作品集
// @Produce  application/json
// @Param    userName path string true "user Name"
// @Success  200
// @Router   /v1/portfolio/{userName} [get]
func (p *PortfolioApi) GetUserPortfolio(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No userName found"})
		return
	}

	portfolio, err := portfolioService.GetUserPortfolio(c, userName)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithData(portfolio, c)
}

// All Portfolio
// @Tags     Portfolio
// @Summary  個人作品集列表
// @Produce  application/json
// @Param    data  body      common.Pagination ture "取得分頁列表"
// @Success  200
// @Router   /v1/portfolio/list [post]
func (p *PortfolioApi) GetAllPortfolio(c *gin.Context) {
	var pageRequest common.Pagination
	err := c.ShouldBindJSON(&pageRequest)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	portfolios, err := portfolioService.GetAllPortfolio(c, pageRequest)
	if err != nil {
		global.EP_LOG.Error("取得列表失敗!", zap.Error(err))
		common.FailWithDetailed(nil, err.Error(), c)
		return
	}
	common.OkWithDetailed(portfolios, "取得列表成功", c)

}
