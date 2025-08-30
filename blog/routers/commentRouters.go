package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/comment")
	{
		userRouters.POST("/add", controller.CommentController{}.AddComment)
		userRouters.POST("/delete", controller.CommentController{}.DelComment)
		userRouters.POST("/postComments", controller.CommentController{}.QueryPostComments)
	}
}
