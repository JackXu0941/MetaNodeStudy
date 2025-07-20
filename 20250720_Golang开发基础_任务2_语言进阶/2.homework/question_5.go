package main

import (
	"fmt"
)

// 第五题 面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// 终端运行    go run main.go Question_5.go

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法
type Shape interface {
	Area()
	Perimeter()
}

// Rectangle 结构体
type Rectangle struct {
	name      string
	area      string
	perimeter string
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle Area:", r.area)
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter:", r.perimeter)
}

// Circle 结构体
type Circle struct {
	name      string
	area      string
	perimeter string
}

func (c *Circle) Area() {
	fmt.Println("Circle Area:", c.area)
}

func (c *Circle) Perimeter() {
	fmt.Println("Circle Perimeter:", c.perimeter)
}

func ShoWeShape(shape Shape) {
	shape.Area()
	shape.Perimeter()
}

func Question_5() {

	rectangle := Rectangle{name: "Rectangle", area: "100", perimeter: "200"}
	circle := Circle{name: "Circle", area: "300", perimeter: "500"}
	fmt.Println("矩形信息:")
	ShoWeShape(&rectangle)
	fmt.Println("圆形信息:")
	ShoWeShape(&circle)

}
