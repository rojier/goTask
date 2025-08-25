package models

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

// 表示配置操作数据库的表名称
func (Student) TableName() string {
	return "learn_student"
}
