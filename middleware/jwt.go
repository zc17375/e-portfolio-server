package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zc17375/e-portfolio-server/model/common"
	"github.com/zc17375/e-portfolio-server/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		token := c.GetHeader("x-token")
		if token == "" {
			common.NoAuth("沒有權限登入", c)
			c.Abort()
			return
		}
		// // 黑名單
		// if jwtService.IsBlacklist(token) {
		// 	common.NoAuth("此帳號沒有權限登入", c)
		// 	utils.ClearToken(c)
		// 	c.Abort()
		// 	return
		// }
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				common.NoAuth("權限已過期，請重新登入", c)
				c.Abort()
				return
			}
			common.NoAuth(err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
