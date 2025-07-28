1 整个项目包括了作业的 1~5 个部分，第六题错误处理与日志记录限于时间有限只做了文件没有配置
2 整个项目都是进过实际运行和测试的 , 可以正常运行
3 项目结构
  - s20250727_Golang开发基础_任务4_个人博客系统后端
    - main.go
    - go.mod
    - go.sum
    - dbfiles
        - middleware
           - auth.go
        - createDB.go
    - handlers
        - commet.go
        - post.go
    -userLoginRegister
        - userLoginRegister.go
    - utils
        - error.go
        - jwt.go
    - README.md

  4 项目依赖包文件 go.sum , go.mod 和项目文件都在项目文件夹下
  5 运行方式 :
    1. 启动服务器 项目文件夹目录> go run main.go
    2. 先运行  第二部分  dbfiles.CreateDB() ,穿件数据库 , 其他部分可先注释
    3. 每次运行测试, 请先 访问服务器 // 127.0.0.1:8080/register 和 // 127.0.0.1:8080/login 完成注册和登录; 如果已经注册过, 请先登录
    4. 登录后,会返回 JWT token, 请将 token 保存起来, 登录后, 请将 token 放在请求头中, 请求头中添加 Authorization: Bearer <token>
    5. 然后请多创建几篇文章, 分别在不同的用户中, 测试文章的权限控制 
    6. 请按照main 函数中, 每个不同的 函数的注释, 测试对应的功能
    7. 测试过程中, 请勿关闭服务器, 否则, token 将会失效, 请重新登录
    8. 测试过程中, 可以在数据库中查看数据, 手动插入数据, 测试功能
 6. 测试代码,建议使用 visual studio code , 配置好 go 语言环境, 然后打开项目, 进行测试