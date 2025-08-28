package middleware

import (
	"blog/tool"
	"fmt"
	"net/http"
	"os"
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
			tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
			fmt.Println("================111111111111")
			fmt.Println(tokenString)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("JWT Parse Error: "+err.Error(), tool.ERROR))
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("userId", claims["userId"])
				c.Set("userName", claims["userName"])
				c.Next()
			} else {
				// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.AbortWithStatusJSON(http.StatusOK, tool.ReponseErrorMsg("invalid token", tool.ERROR))
			}
		}

	}
}
