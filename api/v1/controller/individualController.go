package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/model/request"
	"go.uber.org/zap"
)

type IndividualApi struct{}

// UserTest
// @Tags     User
// @Summary  新增個人作品集資料
// @Produce   application/json
// @Param    data  body      model.Individual true "個人作品集資料"
// @Success  200
// @Router   /v1/individual/create [post]
// @Security ApiKeyAuth
func (i *IndividualApi) CreateIndividual(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
		return
	}
	// 取得使用者ID
	customClaims := claims.(*request.CustomClaims)
	user := customClaims.BaseClaims

	var reciveData model.Individual
	err := c.ShouldBindJSON(&reciveData)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	// 新增資料
	newIndivi, err := individualService.CreateIndividual(c, reciveData, user.UUID.String())
	if err != nil {
		global.EP_LOG.Error(user.UUID.String() + "新增Individual資料失敗", zap.Error(err))
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithDetailed(newIndivi, "新增資料成功", c)
}

// UserTest
// @Tags     User
// @Summary  更新個人作品集資料
// @Produce   application/json
// @Param    data  body      model.Individual true "更新個人作品集資料"
// @Success  200
// @Router   /v1/individual/update [post]
// @Security ApiKeyAuth
func (i *IndividualApi) UpdateIndividual(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
		return
	}
	// 取得使用者ID
	customClaims := claims.(*request.CustomClaims)
	user := customClaims.BaseClaims

	var reciveData model.Individual
	err := c.ShouldBindJSON(&reciveData)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	// 更新資料
	updateIndivi, err := individualService.UpdateIndividual(c, reciveData, user.UUID.String())
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithDetailed(updateIndivi, "更新資料成功", c)

}
