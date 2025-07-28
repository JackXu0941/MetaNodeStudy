package main

import (
	"test4/dbfiles/middleware"
	"test4/handlers"
	"test4/userLoginRegister"

	"github.com/gin-gonic/gin"
)

func main() {

	// 2.数据库设计与模型定义
	// 设计数据库表结构，至少包含以下几个表：
	// users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
	// posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、
	// created_at 、 updated_at 等字段。
	// comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、
	// post_id （关联 posts 表的 id ）、 created_at 等字段。
	// 使用 GORM 定义对应的 Go 模型结构体。
	// 可以在 终端 运行命令: go run main.go createDB.go

	// dbfiles.CreateDB()

	// 3.用户认证与授权
	// 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
	// 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，
	// 后续的需要认证的接口需要验证该 JWT 的有效性。

	r := gin.Default()

	// 注册
	// 127.0.0.1:8080/register  用 json 数据发送 ,{ "u":"2222","p" : "2222","e" : "e2222"}
	r.POST("/register", userLoginRegister.Register)

	// 登录
	// 127.0.0.1:8080/login  用 json 数据发送 ,{ "u":"2222","p" : "2222","e" : "e2222"}
	r.POST("/login", userLoginRegister.Login)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM2OTQ0NjEsImlkIjo5LCJ1c2VybmFtZSI6IjIyMjIifQ.p1MXQdvy2lA5lhyi9T6aNbomGKPnvCO5ZDZ-2mIzwTM

	// 需要认证的路由组
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{

		// 	4 .文章管理功能
		// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
		// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
		// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
		// 实现文章的删除功能，只有文章的作者才能删除自己的文章。

		// 创建文章   实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
		// http: //127.0.0.1:8080/createposts 用 json 数据发送 ,{ "title":"title_1","content" : "content_1","e" : "e2222"}
		authorized.POST("/createposts", handlers.CreatePost)

		// 实现文章的读取功能，支持获取所有文章列表的详细信息。
		// http: //127.0.0.1:8080/getposts
		authorized.POST("/getposts", handlers.GetPosts)

		// // 实现文章的读取功能，支持获取单个文章的详细信息。
		//127.0.0.1:8080/getpost?id=1
		authorized.POST("/getpost", handlers.GetPost)

		// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
		// 127.0.0.1:8080/UpdatePost?id=1 ,用 json 数据发送 ,{ "title":"title_1","content" : "content_1","e" : "e2222"}
		authorized.POST("/UpdatePost", handlers.UpdatePost)

		// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
		// 127.0.0.1:8080/deletePost?id=1
		authorized.POST("/deletePost", handlers.DeletePost)

		// 5 .评论功能
		// 实现评论的创建功能，已认证的用户可以对文章发表评论。
		// 实现评论的读取功能，支持获取某篇文章的所有评论列表。

		// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
		//127.0.0.1:8080/createComment?post_id=3 ,用 json 数据发送 ,{ "content":"content_3"}
		authorized.POST("/createComment", handlers.CreateComment)

		// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
		//127.0.0.1:8080/getComments?post_id=3
		authorized.POST("/getComments", handlers.GetComments)

	}

	// r.POST("/createpost", handlers.CreatePost)

	// 127.0.0.1:8080/getposts
	// r.POST("/getposts", handlers.GetPosts)

	// r.POST("/getpost", handlers.GetPost)

	// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM3MDc0NjYsImlkIjo5LCJ1c2VybmFtZSI6IjIyMjIifQ.p9WYDNi7Q5he_WxplxbwlyVHTbnHyHYpLWBEg0mWhtQ

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
