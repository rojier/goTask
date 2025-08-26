package models

type Comment struct {
	Id      int
	Content string
	PostId  int //外键
}

func (Comment) TableName() string {
	return "comment"
}
