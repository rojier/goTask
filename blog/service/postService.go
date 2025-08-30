package service

import (
	"blog/dao"
	"errors"
)

type PostService struct {
}

/*
添加文章
*/
func (postService PostService) AddPost(post dao.Post) error {

	if result := dao.DB.Create(&post); result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("插入数据失败")
	}

}
func (postService PostService) UserPosts(userId int, postId int) []dao.Post {
	posts := []dao.Post{}
	if postId == 0 {
		dao.DB.Where("user_id=?", userId).Find(&posts)
	} else {
		dao.DB.Where("user_id=? and id = ?", userId, postId).Find(&posts)
	}

	return posts

}
func (postService PostService) UpdatePost(post dao.Post) error {

	var dbPost dao.Post
	if err := dao.DB.Where("user_id = ? and id = ?", post.UserID, post.ID).First(&dbPost).Error; err != nil {
		return errors.New("该用户无该文章权限")

	}
	post.CreateTime = dbPost.CreateTime
	result := dao.DB.Save(&post)
	if result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("更新失败")
	}
}

func (postService PostService) DeletePost(userId int, postId int) error {
	var dbPost dao.Post
	if err := dao.DB.Where("user_id = ? and id = ?", userId, postId).First(&dbPost).Error; err != nil {
		return errors.New("该用户无该文章权限")

	}
	result := dao.DB.Delete(&dbPost)
	if result.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("更新失败")
	}

}
