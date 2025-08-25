package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(localhost:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //开启打印sql
	})
	if err != nil {
		fmt.Println("数据库连接错误: ", err)
	}
	fmt.Println("走了core的init方法")
}
