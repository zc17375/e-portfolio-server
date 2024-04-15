package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zc17375/e-portfolio-server/global"
	"github.com/zc17375/e-portfolio-server/model"
	"github.com/zc17375/e-portfolio-server/model/request"
	"github.com/zc17375/e-portfolio-server/utils"
	"gorm.io/gorm"
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

func (as *AuthService) DeleteUser(uuid uuid.UUID) (bool, error) {
	var user model.User
	err := global.EP_DB.Where("uuid = ?", uuid.String()).First(&user).Error
	if err != nil {
		return false, err
	}

	// 更新Delete date and disable
	user.Disable = 1
	user.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true} 
	// 更新 DeletedAt 欄位
	result := global.EP_DB.Model(&user).Updates(user)
	if result.Error != nil {
		return false, err
	}

	return true, nil
}
