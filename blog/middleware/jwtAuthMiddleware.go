package middleware

import (
	"blog/constant"
	"blog/tool"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/user/register" || path == "/user/login" {
			fmt.Println(".........." + path + " skip JWTAuth .....")
			c.Next() // 继续执行后续的处理器或中间件
		} else {
			// 从请求头中获取Token，常见格式是 Authorization: Bearer <token>
			tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("请求头token为空", tool.ERROR))
				return
			}
			// 解析和验证Token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// 验证签名方法
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(constant.JWT_SECRET), nil
			})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("JWT Parse Error: "+err.Error(), tool.ERROR))
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var userId any = claims["userId"]
				if mId, error := tool.AnyToInt(userId); error == nil {
					c.Set("userId", mId)
					c.Set("userName", claims["userName"])
					c.Next()
				} else {
					c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("获取token的用户id类型错误", tool.ERROR))
				}

			} else {
				// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("token无效或过期", tool.ERROR))
			}
		}

	}
}
