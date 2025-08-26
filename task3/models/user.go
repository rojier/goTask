package models

//foreignKey外键  如果是表名称加上Id的话默认也可以不配置   如果不是我们需要通过foreignKey配置外键
//references表示的是主键    默认就是Id   如果是Id的话可以不配置
type User struct {
	Id      int
	Name    string
	PostNum int
	Post    []Post `gorm:"foreignKey:UserId;references:Id"`
}

func (User) TableName() string {
	return "user"
}
