package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model/request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.EP_CONFIG.Jwt.SigningKey),
	}
}


func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.EP_CONFIG.Jwt.BufferTime)
	ep, _ := ParseDuration(global.EP_CONFIG.Jwt.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 判斷刷新令牌時間，如果未超過此時間則產生新的token
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"EPS"},                   // 受眾
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.EP_CONFIG.Jwt.Issuer,              // 签名的发行者
		},
	}
	return claims
}

// 生成token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}


// 解析Token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	// 檢查錯誤類型
    if err != nil {
        switch err.Error() {
        case "token contains an invalid number of segments":
            fmt.Println("Token 格式錯誤")
			return nil, TokenMalformed
        case "token is expired":
            fmt.Println("Token 已過期")
			return nil, TokenExpired
        case "token is not valid yet":
            fmt.Println("Token 還未生效")
			return nil, TokenNotValidYet
        case "crypto/ecdsa: verification error":
            fmt.Println("Token 簽名無效")
			return nil, TokenInvalid
        default:
            fmt.Println("其他錯誤:", err)
			return nil, err
        }
    }
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
