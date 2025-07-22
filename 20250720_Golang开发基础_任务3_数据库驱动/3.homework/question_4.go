package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// 第四题 Sqlx入门
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
// 并将结果映射到 Book 结构体切片中，确保类型安全。
// 终端运行    go run main.go Question_4.go

type books struct {
	ID     uint `gorm:"primarykey"`
	Title  string
	Author string
	Price  float64
}

var db_1 *sqlx.DB

func Question_4() {

	//1 .使用 sqlx 连接数据库
	mDb, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True") //连接操作，mDb是返回的数据库实例
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}

	db_1 = mDb

	// 2. 使用 SQL 语句创建表 employees 的 语句
	// createTableSQL_books := `
	// CREATE TABLE IF NOT EXISTS books (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	title VARCHAR(255) NOT NULL,
	// 	author VARCHAR(255),
	// 	price DECIMAL(10, 2)
	// );
	// `

	// 3. 执行创建表的 SQL 语句
	// _, err = db_1.Exec(createTableSQL_books)
	// if err != nil {
	// 	fmt.Println("创建表失败:", err)
	// 	return
	// }
	// fmt.Println("表创建成功或已存在")

	// 4. 向 employees 插入测试数据
	// sqlStr := "insert into books(title, author,price) values (?,?,?)"
	// ret, err := db_1.Exec(sqlStr, "字典1", "刘一", 35)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err = db_1.Exec(sqlStr, "小说1", "陈二", 25)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err = db_1.Exec(sqlStr, "诗词1", "张三", 55)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// ret, err = db_1.Exec(sqlStr, "漫画", "李四", 50)
	// if err != nil {
	// 	fmt.Printf("insert failed, err:%v\n", err)
	// 	return
	// }
	// print(ret.RowsAffected())

	// 5.编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
	// 并将结果映射到 Book 结构体切片中，确保类型安全。
	var querySQL_1 = "select * from books where price >=?"
	var books_rows []books
	err1 := db_1.Select(&books_rows, querySQL_1, 50) //很经典的sql语句
	if err1 != nil {
		fmt.Println("查询异常, ", err1)
		return
	}
	for _, v := range books_rows {
		fmt.Println(v.ID, v.Title, v.Author, v.Price)
	}

}
