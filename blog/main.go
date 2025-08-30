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
	r.Use(middleware.LoggerMidderware(), middleware.JWTAuth())
	//注册路由
	routers.UserRoutersInit(r)
	routers.PostRoutersInit(r)
	routers.CommentRoutersInit(r)
	r.Run(":8081")
}
