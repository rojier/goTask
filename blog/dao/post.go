package dao

import "time"

type Post struct {
	ID         uint   `gorm:"primarykey"`
	Title      string `gorm:"not null" json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	UserID     uint   `gorm:"not null"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (Post) TableName() string {
	return "post"
}
