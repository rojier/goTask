package dao

type User struct {
	Id       int
	UserName string `json:"username"`
	PassWord string `json:"passWord"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
