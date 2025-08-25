package model

// 用户表
type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Posts     []Post `gorm:"foregnKey:UserID"`
	PostCount int
}

// 评论表
type Comment struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Content string
	UserID  uint `gorm:"not null"` // 外键
	PostID  uint `gorm:"not null"`
}

// 文章
type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentStatus string    // 评论状态
	CommentCount  int       // 评论数量
}

// package main

// import (
//     "gorm.io/driver/sqlite"
//     "gorm.io/gorm"
// )

// type User struct {
//     gorm.Model
//     Username string `gorm:"unique;not null"`
//     Password string `gorm:"not null"`
//     Email    string `gorm:"unique;not null"`
// }

// type Post struct {
//     gorm.Model
//     Title   string `gorm:"not null"`
//     Content string `gorm:"not null"`
//     UserID  uint
//     User    User
// }

// type Comment struct {
//     gorm.Model
//     Content string `gorm:"not null"`
//     UserID  uint
//     User    User
//     PostID  uint
//     Post    Post
// }

// func main() {
//     db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }

//     // 自动迁移模型
//     db.AutoMigrate(&User{}, &Post{}, &Comment{})
// }
