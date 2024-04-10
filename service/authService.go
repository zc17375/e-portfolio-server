package service

import (
	"errors"
	"fmt"

	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
)

type AuthService struct{}

func (as *AuthService) Login(u *model.User) (*model.User, error) {
	if nil == global.EP_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user model.User
	err := global.EP_DB.Where("email = ? OR phone = ?", u.Email, u.Email).First(&user).Error
	if err == nil {
		// if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		// 	return nil, errors.New("帳號或密碼錯誤")
		// }
		if ok := (u.Password == user.Password); !ok {
			return nil, errors.New("帳號或密碼錯誤")
		}
	}
	return &user, err
}
