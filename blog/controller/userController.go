package controller

import (
	"blog/dao"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// 用于绑定 JSON 参数的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" `
}

func (userController UserController) Register(c *gin.Context) {
	var loginReq LoginRequest
	// 使用 ShouldBindJSON 绑定 JSON，并处理错误
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		userController.RspError(c)
		return
	}
	user := dao.User{
		UserName: loginReq.Username,
		PassWord: loginReq.Password,
		Email:    loginReq.Email,
	}
	err := service.UserService{}.Regitser(user)
	userController.RspCommon(c, err)
}

func (userController UserController) Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		userController.RspError(c)
		return
	}
	user := dao.User{
		UserName: loginReq.Username,
		PassWord: loginReq.Password,
	}

	token, err := service.UserService{}.Login(user)
	if err != nil {
		userController.RspEMsg(c, err.Error())
	} else {
		//将令牌添加到请求头中``
		// c.Request.Header.Set("Authorization", "Bearer "+token)
		userController.RspSuccess(c, token)
	}

}
