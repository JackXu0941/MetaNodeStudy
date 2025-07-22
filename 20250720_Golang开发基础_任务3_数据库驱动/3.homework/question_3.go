package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// 第三题 Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、
// department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
// 终端运行    go run main.go Question_3.go

type employees struct {
	ID         uint `gorm:"primarykey"`
	Name       string
	Department string
	Salary     float64
}

var db *sqlx.DB

func Question_3() {

	//1 .使用 sqlx 连接数据库
	mDb, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True") //连接操作，mDb是返回的数据库实例
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}

	db = mDb

	// 2. 使用 SQL 语句创建表 employees 的 语句
	// createTableSQL := `
	// CREATE TABLE IF NOT EXISTS employees (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	name VARCHAR(255) NOT NULL,
	// 	department VARCHAR(255),
	// 	salary DECIMAL(10, 2)
	// );
	// `

	// 3. 执行创建表的 SQL 语句
	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	fmt.Println("创建表失败:", err)
	// 	return
	// }
	// fmt.Println("表创建成功或已存在")

	// 4. 向 employees 插入测试数据
	// sqlStr := "insert into employees(name, department,salary) values (?,?,?)"
	// ret, err := db.Exec(sqlStr, "王五", "技术部", 25000)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err := db.Exec(sqlStr, "刘一", "技术部", 35000)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err = db.Exec(sqlStr, "陈二", "销售部", 35000)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err = db.Exec(sqlStr, "张三", "财务部", 35000)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// 5 .编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
	// 并将结果映射到一个自定义的 Employee 结构体切片中。
	// var querySQL_1 = "select * from employees where department=?"
	// var employees_rows []employees
	// err1 := db.Select(&employees_rows, querySQL_1, "技术部") //很经典的sql语句
	// if err1 != nil {
	// 	fmt.Println("查询异常, ", err1)
	// 	return
	// }
	// for _, v := range employees_rows {
	// 	fmt.Println(v.ID, v.Name, v.Department, v.Salary)
	// }

	//6 .编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，
	// 并将结果映射到一个 Employee 结构体中。
	var querySQL_2 = "select * from employees order by salary desc limit 1"
	var employees_row employees
	err = db.QueryRow(querySQL_2).Scan(&employees_row.ID, &employees_row.Name, &employees_row.Department, &employees_row.Salary)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println(employees_row.ID, employees_row.Name, employees_row.Department, employees_row.Salary)

}
