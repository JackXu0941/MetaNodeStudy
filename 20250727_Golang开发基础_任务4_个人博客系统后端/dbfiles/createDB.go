package dbfiles

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"u" uri:"u" form:"u" binding:"required"`
	Password string `gorm:"not null" json:"p" uri:"p" form:"p" binding:"required"`
	Email    string `gorm:"unique;not null" json:"e" uri:"e" form:"e" `
}

// type Post struct {
// 	gorm.Model
// 	Title   string `gorm:"not null"`
// 	Content string `gorm:"not null"`
// 	UserID  uint
// 	User    User
// }

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title" binding:"required"`
	Content   string    `gorm:"type:text;not null" json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"` // 关联用户信息
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
}

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func CreateDB() {
	// db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	db := InitDB()
	// 自动迁移模型
	// db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.AutoMigrate(&Comment{})
}
