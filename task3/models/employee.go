package models

type Employee struct {
	Id   int
	Name string
	//部门
	Department string
	//工资
	Salary int
}

func (Employee) TableName() string {
	return "employee"
}
