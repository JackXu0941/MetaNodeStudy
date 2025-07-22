package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID    uint `gorm:"primary_key"`
	Name  string
	Age   int
	Grade string
}

// 第一题 SQL语句练习
// 1.假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
// age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 2.编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 3.编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 4.编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 5.编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
// 终端运行    go run main.go Question_1.go
func Question_1(db *gorm.DB) {

	//1 创建表
	db.AutoMigrate(&User{})

	// //2 插入数据
	user := User{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&user) // 通过数据的指针来创建插入,否则不会插入
	fmt.Println(result.RowsAffected)

	// //3 根据条件查询数据
	var user1 User
	db.Debug().First(&user1, "age > ?", "18")
	fmt.Println(user1)

	// //4 根据条件更新数据
	db.Debug().First(&user1, "name = ?", "张三")
	fmt.Println(user1)
	user1.Grade = "四年级"
	db.Save(&user1)
	db.Debug().First(&user1, "name = ?", "张三")
	fmt.Println(user1)

	//5 根据条件删除数据
	// var user2 User
	// db.Where("age < ?", "22").Delete(user2)
	// fmt.Println(user2)

}
