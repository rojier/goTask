package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	{
		userRouters.POST("/register", controller.UserController{}.Register)

	}
}
