package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/model/request"
	"github.com/zc17375/e-portfolio-server/model/response"
	"github.com/zc17375/e-portfolio-server/utils"
	"go.uber.org/zap"
)

type AuthApi struct{}

// Login
// @Tags     Auth
// @Summary  會員登入
// @Produce   application/json
// @Param    data  body      request.Login   true  "信箱或手機號碼, 密碼"
// @Success  200   {object}  common.Response{data=response.LoginResponse,msg=string}  "返回使用者資訊,token,過期時間"
// @Router   /v1/auth/login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	user, err := authService.Login(&l)
	if err != nil {
		global.EP_LOG.Error("登入失敗! 帳號不存在或密碼錯誤!", zap.Error(err))
		common.FailWithMessage("帳號不存在或密碼錯誤", c)
		return
	}
	if user.Disable == 1 {
		global.EP_LOG.Error("登入失敗! 此會員已註銷")
		common.FailWithMessage("此會員已註銷", c)
		return
	}
	// 取得Token並返回結果
	a.GenerateToken(c, *user)
}

// 生成Auth Token
func (a *AuthApi) GenerateToken(c *gin.Context, user model.User) {
	j := &utils.JWT{SigningKey: []byte(global.EP_CONFIG.Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
	})
	
	// 生成Token
	token, err := j.CreateToken(claims)
	if err != nil {
		global.EP_LOG.Error("取得JWT Token失敗!", zap.Error(err))
		common.FailWithMessage("取得JWT Token失敗", c)
		return
	}

	// 將token存到DB
    user.JwtToken = token
	user.LastLoginAt = time.Now()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
    if err := global.EP_DB.Save(&user).Error; err != nil {
        global.EP_LOG.Error("DB新增Token失敗!", zap.Error(err))
		common.FailWithMessage("新增JWT Token失敗", c)
		return
    }

	common.OkWithDetailed(response.LoginResponse{
		User:      user,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登入成功", c)

}

// Register
// @Tags     Auth
// @Summary  會員註冊
// @Produce   application/json
// @Param    data  body      request.Register                                            true  "使用者名稱, 密碼, 暱稱, 頭貼網址, 手機號碼, 信箱"
// @Success  200   {object}  common.Response{data=response.UserResponse,msg=string}  "會員註冊,返回會員註冊訊息"
// @Router   /v1/auth/register [post]
func (a *AuthApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	user := &model.User{
		Username: r.Username, 
		NickName: r.NickName, 
		Password: r.Password, 
		HeaderImg: r.HeaderImg,
		Phone: r.Phone, 
		Email: r.Email,
	}
	userReturn, err := authService.Register(*user)
	if err != nil {
		global.EP_LOG.Error("註冊失敗!", zap.Error(err))
		common.FailWithDetailed(nil, "註冊會員失敗:此信箱或手機號碼已被註冊", c)
		return
	}
	common.OkWithDetailed(response.UserResponse{User: userReturn}, "註冊會員成功", c)

}