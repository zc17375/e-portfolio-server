package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/model/common"
)

type UserApi struct{}

// UserTest
// @Tags     User
// @Summary  token測試
// @Produce   application/json
// @Success  200
// @Router   /v1/user/portfolio [get]
// @Security ApiKeyAuth
func (u *UserApi) Portfolio(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
		return
	}
	common.OkWithData(claims, c)
}
