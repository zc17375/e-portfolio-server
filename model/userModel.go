package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zc17375/e-portfolio-server/global"
	"gorm.io/gorm"
)

type User struct {
	global.EP_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:使用者UUID"`   // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:使用者名稱"` // 用户登录名
	Password    string    `json:"-"  gorm:"comment:使用者密碼"`             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"comment:暱稱"`
	HeaderImg   string    `json:"headerImg" gorm:"default:https://img.jpg;comment:使用者頭像"` // 用户头像
	Phone       string    `json:"phone"  gorm:"comment:手機號碼"`                             // 手機號碼
	Email       string    `json:"email"  gorm:"comment:電子郵箱"`                             // 電子信箱
	Disable     int       `json:"disable" gorm:"default:0;comment:是否註銷:0=註銷"`             //帳號是否註銷
	JwtToken    string    `json:"jwt_token" gorm:"comment:JWT Token;type:text"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"default:null;comment:上次登入時間"`
}


func (model *User) IsUserExist(u User) bool {
	// 判断手機或信箱是否注册
	return !errors.Is(global.EP_DB.Where("phone = ? OR email = ?", u.Phone, u.Email).First(model).Error, gorm.ErrRecordNotFound)
}