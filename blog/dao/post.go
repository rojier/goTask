package dao

import "time"

type Post struct {
	ID         uint   `gorm:"primarykey"`
	Title      string `gorm:"not null"`
	Content    string
	UserID     uint `gorm:"not null"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (Post) TableName() string {
	return "post"
}
