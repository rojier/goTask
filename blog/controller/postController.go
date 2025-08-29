package controller

import (
	"blog/dao"
	"blog/service"
	"time"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	BaseController
}

// 用于绑定 JSON 参数的结构体
type AddPostRequest struct {
	Content string `json:"content" binding:"required"`
	Title   string `json:"title" binding:"required"`
}
type UserPostRequest struct {
	PostID int `json:"postID"`
}

/*
添加文章
*/
func (postController PostController) AddPost(c *gin.Context) {

	var addPostRequest AddPostRequest
	if err := c.ShouldBindJSON(&addPostRequest); err != nil {
		postController.RspEMsg(c, "请求参数错误")
		return
	}
	userId := c.GetInt("userId")
	post := dao.Post{
		Content:    addPostRequest.Content,
		Title:      addPostRequest.Title,
		UserID:     uint(userId),
		CreateTime: time.Now(), // 不设置会插入空值 0000-00-00 00:00:00
		UpdateTime: time.Now(),
	}
	err := service.PostService{}.AddPost(post)
	postController.RspCommon(c, err)
}

/*
获取用户的文章列表
*/
func (postController PostController) UserPosts(c *gin.Context) {
	var userPostRequest UserPostRequest
	if err := c.ShouldBindJSON(&userPostRequest); err != nil {
		postController.RspEMsg(c, "请求参数错误")
		return
	}
	posts := service.PostService{}.UserPosts(c.GetInt("userId"), userPostRequest.PostID)
	postController.RspSuccess(c, posts)
}
