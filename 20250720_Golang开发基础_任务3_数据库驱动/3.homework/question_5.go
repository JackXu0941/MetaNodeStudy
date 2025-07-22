package main

import (
	"gorm.io/gorm"
)

// 第五题 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，
// 则更新文章的评论状态为 "无评论"。
// 终端运行    go run main.go Question_5.go

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法
type User_1 struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Posts       []Post `gorm:"foreignkey:UserID"`
	Posts_count int
}

type Post struct {
	ID       uint `gorm:"primary_key"`
	Content  string
	UserID   uint
	Comments []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	ID      uint `gorm:"primary_key"`
	Content string
	PostID  uint
}

func CreateUser(db *gorm.DB) {
	user := User_1{
		Name: "张三",
		Posts: []Post{
			{
				Content: "Go语言入门教程",
				Comments: []Comment{
					{Content: "写的真好！"},
					{Content: "有帮助，谢谢分享"},
				},
			},
			{
				Content: "GORM使用指南",
				Comments: []Comment{
					{Content: "非常详细的文档"},
					{Content: "示例代码很实用"},
					{Content: "收藏了"},
				},
			},
		},
	}

	// 插入用户及其关联的文章和评论
	db.Create(&user)
}

func Question_5(db *gorm.DB) {

	// 题目1：模型定义
	// 创建表
	db.AutoMigrate(&User_1{}, &Post{}, &Comment{})

	// 插入用户及其关联的文章和评论 的记录
	CreateUser(db)

	//题目2：关联查询
	//1 .使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// var user2 User_1
	// db.Debug().Preload("Posts.Comments").First(&user2, "id = ?", 1)
	// fmt.Println("User:", user2.ID, user2.Name)
	// for _, post := range user2.Posts {
	// 	fmt.Printf("Post ID: %d, Content: %s\n", post.ID, post.Content)
	// 	for _, comment := range post.Comments {
	// 		fmt.Printf("  Comment ID: %d, Content: %s\n", comment.ID, comment.Content)
	// 	}
	// }

	//2 .使用Gorm查询评论数量最多的文章信息

	// var post Post
	// db.Debug().
	// 	Table("posts").
	// 	Select("posts.*").
	// 	Joins("LEFT JOIN comments ON comments.post_id = posts.id").
	// 	Group("posts.id").
	// 	Order("COUNT(comments.id) DESC").
	// 	Limit(1).
	// 	Scan(&post)
	// // 获取该文章的所有评论
	// db.Debug().Model(&post).Preload("Comments").Find(&post)

	// fmt.Printf("Post ID: %d, Content: %s, Comment Count: %d\n", post.ID, post.Content, len(post.Comments))

}

// 题目3：钩子函数
// 1 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User_1
	tx.First(&user, p.UserID)
	tx.Model(&user).UpdateColumn("posts_count", gorm.Expr("posts_count + ?", 1))
	return
}

// 题目3：钩子函数
// 2 为Comment模型添加钩子函数，在评论删除时检查文章的评论数量：
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	tx.First(&post, c.PostID)

	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	if count == 0 {
		tx.Model(&post).Update("comment_status", "无评论")
	}

	return
}
