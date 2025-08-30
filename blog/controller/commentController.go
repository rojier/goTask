package controller

import (
	"blog/dao"
	"blog/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	BaseController
}

// 用于绑定 JSON 参数的结构体
type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostId  uint   `json:"postId" binding:"required"`
}
type DelCommentRequest struct {
	CommentId uint `json:"commentId" binding:"required"`
}
type PostCommentsRequest struct {
	PostId uint `json:"postId" binding:"required"`
}

/*
添加文章
*/
func (commentController CommentController) AddComment(c *gin.Context) {

	var addCommentRequest AddCommentRequest
	if err := c.ShouldBindJSON(&addCommentRequest); err != nil {
		commentController.RspEMsg(c, "请求参数错误")
		return
	}

	userId := c.GetInt("userId")
	comment := dao.Comment{
		Content:    addCommentRequest.Content,
		PostId:     addCommentRequest.PostId,
		UserID:     uint(userId),
		CreateTime: time.Now(), // 不设置会插入空值 0000-00-00 00:00:00
		UpdateTime: time.Now(),
	}
	err := service.CommentService{}.AddComment(comment)
	commentController.RspCommon(c, err)
}

/*
删除文章
*/
func (commentController CommentController) DelComment(c *gin.Context) {

	var delCommentRequest DelCommentRequest
	if err := c.ShouldBindJSON(&delCommentRequest); err != nil {
		commentController.RspEMsg(c, "请求参数错误")
		return
	}
	userId := c.GetInt("userId")
	err := service.CommentService{}.DelComment(uint(userId), delCommentRequest.CommentId)
	commentController.RspCommon(c, err)

}

/*
查询文章的评论
*/
func (commentController CommentController) QueryPostComments(c *gin.Context) {
	var postCommentsRequest PostCommentsRequest
	if err := c.ShouldBindJSON(&postCommentsRequest); err != nil {
		commentController.RspEMsg(c, "请求参数错误")
		return
	}
	comments := service.CommentService{}.QueryPostComments(postCommentsRequest.PostId)
	commentController.RspSuccess(c, comments)
}
