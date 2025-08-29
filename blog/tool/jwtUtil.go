package tool

import (
	"blog/constant"
	"blog/dao"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
JWT SecretKey
密钥长度至少 32 字节（256 位），以应对暴力破解
*/

func GenerateSecretKey() (string, error) {
	key := make([]byte, 32) // 256位
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

/*
生成JWT Token的函数
*/
func GenerateToken(storedUser dao.User) (string, error) {

	// 使用HS256算法创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   storedUser.Id,
		"userName": storedUser.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	// 使用密钥签名Token
	return token.SignedString([]byte(constant.JWT_SECRET))
}

func TestJwt() {
	storedUser := dao.User{
		UserName: "测试",
		Id:       1,
	}
	// 生成 token
	tokenString, _ := GenerateToken(storedUser)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.JWT_SECRET), nil
	})

	fmt.Println("token:", token)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("验证通过")
	}
}
