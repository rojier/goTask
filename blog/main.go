package main

import (
	"blog/middleware"
	"blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//添加中间件
	r.Use(middleware.LoggerMidderware())
	//注册路由
	routers.UserRoutersInit(r)
	r.Run(":8081")
}
