package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/request"
	"github.com/zc17375/e-portfolio-server/utils"
)

type AuthService struct{}

func (as *AuthService) Login(l *request.Login) (*model.User, error) {
	if nil == global.EP_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user model.User
	err := global.EP_DB.Where("email = ? OR phone = ?", l.Account, l.Account).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(l.Password, user.Password); !ok {
			return nil, errors.New("帳號或密碼錯誤")
		}
	}
	return &user, err
}

func (as *AuthService) Register(u model.User) (model.User, error) {
	var user model.User
	// 判断手機或信箱是否已註冊
	// if !errors.Is(global.EP_DB.Where("phone = ? OR email = ?", u.Phone, u.Email).First(&user).Error, gorm.ErrRecordNotFound) { 
	// 	return user, errors.New("此信箱或手機號碼已被註冊")
	// }

	if user.IsUserExist(u) {
		return user, errors.New("此信箱或手機號碼已被註冊")
	}

	// 新增uuid並且密碼加密
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err := global.EP_DB.Create(&u).Error
	return u, err
}
