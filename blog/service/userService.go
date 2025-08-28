package service

import (
	"blog/dao"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

/*
注册用户
*/
func (userService UserService) Regitser(user dao.User) error {
	users := []dao.User{}
	dao.DB.Where("user_name=?", user.UserName).Find(&users)
	if len(users) > 0 {
		return errors.New("该用户名已注册")
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("加密密码错误")
	}
	user.PassWord = string(hashedPassword)
	result := dao.DB.Create(&user)
	if result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("插入数据失败")
	}

}
