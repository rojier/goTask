package main

import (
	"blog/middleware"
	"blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// //添加中间件
	r.Use(middleware.LoggerMidderware(), middleware.JWTAuth())
	// //注册路由
	routers.UserRoutersInit(r)
	r.Run(":8081")

	// var jwtSecret []byte
	// jwtSecret, _ = tool.GenerateSecretKey()

	// if err != nil {
	// 	fmt.Println("Failed to generate secret key:", err)
	// }
	// 打印 base64 编码的密钥（在实际应用中不要这样做，这里只是为了演示）
	// fmt.Println(jwtSecret)
	// fmt.Printf("Generated SecretKey (base64): %s\n", base64.StdEncoding.EncodeToString(jwtSecret))

}
