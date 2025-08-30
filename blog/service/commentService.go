package service

import (
	"blog/dao"
	"errors"
)

type CommentService struct {
}

/*
添加评论
*/
func (commentService CommentService) AddComment(comment dao.Comment) error {

	var dbPost dao.Post
	if err := dao.DB.Where("id = ?", comment.PostId).First(&dbPost).Error; err != nil {
		return errors.New("评论文章为空")

	}
	if result := dao.DB.Create(&comment); result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("插入数据失败")
	}
}

/*
删除评论
*/
func (commentService CommentService) DelComment(userId uint, commentId uint) error {
	var dbComment dao.Comment
	if err := dao.DB.Where("user_id = ? and id = ?", userId, commentId).First(&dbComment).Error; err != nil {
		return errors.New("该用户无该评论权限")
	}
	comcomment := dao.Comment{
		ID: commentId,
	}
	result := dao.DB.Delete(&comcomment)
	if result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("更新失败")
	}

}

/*
查询文章的评论
*/
func (comcomment CommentService) QueryPostComments(postId uint) []dao.Comment {
	comments := []dao.Comment{}
	dao.DB.Where("post_id=?", postId).Find(&comments)
	return comments
}
