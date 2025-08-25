package interface_test

import "fmt"

type Shape interface {
	Area()
	// Perimeter()
}

type Rectangle struct {
	Name string
}

func (s Rectangle) Area() {
	fmt.Println(s.Name, ":Area")
}

type Circle struct {
	Name string
}

func (c Circle) Area() {
	fmt.Println(c.Name, ":Area")
}

func TestInterface() {
	re := Rectangle{
		Name: "Ret",
	}
	var s Shape = re
	s.Area()
	c := Circle{
		Name: "Circle",
	}
	var s1 Shape = c
	s1.Area()
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeId int
	Person     Person
}

func (p *Employee) SetName(name string) {
	p.Person.Name = name
}

func (p Employee) PrintEm() {
	fmt.Println("name: ", p.Person.Name, "age: ", p.Person.Age, "Id: ", p.EmployeeId)
}

func TestStruct() {
	p := Employee{
		Person: Person{
			Name: "员工",
			Age:  18,
		},

		EmployeeId: 1,
	}
	p.PrintEm()
	p.SetName("员工修改")
	p.PrintEm()
}
