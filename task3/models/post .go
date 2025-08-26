package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	Id      int
	Title   string
	UserId  int       //外键
	Comment []Comment `gorm:"foreignKey:PostId;references:Id"`
}

func (t *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Println("添加文章的钩子函数,hook AfterCreate")
	user := User{Id: t.UserId}
	tx.Find(&user)
	user.PostNum = user.PostNum + 1
	tx.Save(&user)
	return nil
}
func (Post) TableName() string {
	return "post"
}
