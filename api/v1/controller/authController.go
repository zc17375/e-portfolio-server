package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/model/request"
	"go.uber.org/zap"
)

type AuthApi struct{}

// Login
// @Tags     Auth
// @Summary  會員登入
// @Produce   application/json
// @Param    data  body      request.Login   true  "使用者名稱, 密碼, 驗證碼"
// @Success  200   {object}  common.Response{data=response.LoginResponse,msg=string}  "返回使用者資訊,token,過期時間"
// @Router   /v1/auth/login [post]
func (u *AuthApi) Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	rUser := &model.User{Username: l.Username, Password: l.Password}
	user, err := authService.Login(rUser)
	if err != nil {
		global.EP_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		common.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.EP_LOG.Error("登陆失败! 用户被禁止登录!")
		common.FailWithMessage("用户被禁止登录", c)
		return
	}
	// 生成Token
	// u.TokenNext(c, *user)
	common.FailWithMessage(global.Translate("sys_user.vCodeErr"), c)


}
