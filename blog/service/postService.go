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
