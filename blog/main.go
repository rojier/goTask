package main

import (
	"blog/middleware"
	"blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	middleware.InitLogger()
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.Use(gin.Recovery())
	//添加中间件
	r.Use(middleware.LoggerMidderware(), middleware.JWTAuth())
	//注册路由
	routers.UserRoutersInit(r)
	routers.PostRoutersInit(r)
	routers.CommentRoutersInit(r)
	r.Run(":8081")
}
