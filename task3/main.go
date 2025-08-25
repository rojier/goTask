package main

import (
	"fmt"
	"task3/models"
	"time"
)

/*
测试插入学生表数据
*/
func testInsertStudent() {

	template := "2006-01-02 15:04:05"
	var now = time.Now().Format(template)
	student := models.Student{
		Name:  "张三" + now,
		Age:   18,
		Grade: "高三",
	}
	models.DB.Create(&student)
	fmt.Println("创建学生数据成功")
}
func queryByAge() {
	studentList := []models.Student{}
	models.DB.Where("age>10").Find(&studentList)
	fmt.Println("查出的数据:", studentList)
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func updateStudent() {
	//更新单个列
	student := models.Student{}
	models.DB.Model(&student).Where("name like ?", "张三%").Update("grade", "四年级")

}
func deleteStudent() {
	student := models.Student{}
	models.DB.Where("age < ?", 20).Delete(&student)
	fmt.Println("删除成功")
}
func main() {
	deleteStudent()
	testInsertStudent()
	updateStudent()
	queryByAge()
}
