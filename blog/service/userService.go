package service

import (
	"blog/constant"
	"blog/dao"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	// 加密密码 该方法对同一字符串加密后每次 结果不一样
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

/*
用户登录
*/
func (userService UserService) Login(user dao.User) (string, error) {
	// users := []dao.User{}
	// dao.DB.Where("user_name=?", user.UserName).Find(&users)
	// if len(users) <= 0 {
	// 	return errors.New("用户未注册")
	// }
	// storedUser := users[0]

	var storedUser dao.User
	if err := dao.DB.Where("user_name = ?", user.UserName).First(&storedUser).Error; err != nil {
		return "", errors.New("用户未注册")

	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.PassWord), []byte(user.PassWord)); err != nil {
		return "", errors.New("密码错误")
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   storedUser.Id,
		"userName": storedUser.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	if mtokenString, err := token.SignedString([]byte(constant.JWT_SECRET)); err != nil {
		return "", errors.New("Failed to generate token")
	} else {
		return mtokenString, nil
	}

}
