package model

import (
	"github.com/zc17375/e-portfolio-server/global"
)

type JwtBlacklist struct {
	global.EP_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
