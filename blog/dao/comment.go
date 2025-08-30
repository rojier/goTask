package dao

import "time"

type Comment struct {
	ID         uint   `gorm:"primarykey"`
	Content    string `json:"content" binding:"required"`
	UserID     uint   `gorm:"not null"`
	PostId     uint   `gorm:"not null"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (Comment) TableName() string {
	return "comment"
}
