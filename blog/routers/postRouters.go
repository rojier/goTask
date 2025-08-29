package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func PostRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/post")
	{
		userRouters.POST("/add", controller.PostController{}.AddPost)
		userRouters.POST("/userPosts", controller.PostController{}.UserPosts)

	}
}
