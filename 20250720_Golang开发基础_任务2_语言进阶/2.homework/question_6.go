package main

import (
	"fmt"
)

// 第六题 面向对象
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，
// 输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体
type Employee struct {
	person     Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	person := e.person
	fmt.Println("EmployeeID:", e.EmployeeID, "Name:", person.Name, "Age:", person.Age)

}

func Question_6() {

	employee := Employee{Person{"张三", 18}, "123456"}
	employee.PrintInfo()

}
