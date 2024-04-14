package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/model/request"
)

type IndividualApi struct{}

// UserTest
// @Tags     User
// @Summary  新增個人經歷資訊
// @Produce   application/json
// @Param    data  body      model.Individual true "個人經歷"
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
	// common.OkWithData(claims, c)

	var reciveData model.Individual
	err := c.ShouldBindJSON(&reciveData)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	// 新增資料
	newIndivi, err := individualService.CreateIndividual(c, reciveData, user.UUID.String())
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(newIndivi)
	common.OkWithDetailed(newIndivi, "新增資料成功", c)
}
