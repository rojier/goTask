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
	// deleteStudent()
	// testInsertStudent()
	// updateStudent()
	// queryByAge()
	accountTest()
}

/*
转账
*/
func accountTest() {
	tx := models.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("转账出错，回滚....")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println("事务初始化错误: ", err)
		return
	}
	//张三账户减去100
	a1 := models.Account{
		Id: 1,
	}
	tx.Find(&a1)
	a1.Balance = a1.Balance - 100
	if err := tx.Save(&a1).Error; err != nil {
		fmt.Println("张三转账出错，回滚....")
		tx.Rollback()
	}
	// panic("测试一个错误")
	a2 := models.Account{
		Id: 2,
	}
	//李四账号增加100
	tx.Find(&a2)
	a2.Balance = a2.Balance + 100
	if err := tx.Save(&a2).Error; err != nil {
		fmt.Println("李四转入出错，回滚....")
		tx.Rollback()
	}
	//记录
	t1 := models.Transaction{
		FromBalanceId: 1,
		ToBalanceId:   2,
		Amount:        100,
	}
	if err := tx.Save(&t1).Error; err != nil {
		fmt.Println("记录出错，回滚....")
		tx.Rollback()
	}
	tx.Commit()
	fmt.Println("转账完成提交....")

}
