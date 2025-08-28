package controller

import (
	"blog/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c BaseController) RspSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, tool.ReponseSuccess(data))
}

func (c BaseController) RspEMsgCode(ctx *gin.Context, errorCode int, errorMsg string) {
	ctx.JSON(http.StatusOK, tool.ReponseErrorMsg(errorMsg, errorCode))
}
func (c BaseController) RspEMsg(ctx *gin.Context, errorMsg string) {
	ctx.JSON(http.StatusOK, tool.ReponseErrorMsg(errorMsg, tool.ERROR))
}
func (c BaseController) RspError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tool.ReponseError())
}
func (b BaseController) RspCommonError(ctx *gin.Context, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, tool.ReponseSuccess(nil))
	} else {
		ctx.JSON(http.StatusOK, tool.ReponseErrorMsg(err.Error(), tool.ERROR))
	}
}
