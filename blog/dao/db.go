package dao

import (
	"blog/constant"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	DB, err = gorm.Open(mysql.Open(constant.DB_PATH), &gorm.Config{
		QueryFields: true, //开启打印
	})
	if err != nil {
		fmt.Println("数据库连接错误:", err)
	}
	fmt.Println("走了db的init方法")
}
