package main

import (
	"fmt"

	"gorm.io/gorm"
)

/**

进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中
User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
**/

// 用户表
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
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

func gormAdvanceInitialTable(db *gorm.DB) {
	if err := db.Debug().AutoMigrate(&User{}, &Comment{}, &Post{}); err != nil {
		panic(err)
	}

	var users = []User{
		{
			ID:   1,
			Name: "张三",
		},
		{
			ID:   2,
			Name: "李四",
		},
	}

	var posts = []Post{
		{
			ID:      1,
			Title:   "post 1",
			Content: "this is post 5",
			UserID:  1,
		},
		{
			ID:      2,
			Title:   "post 2",
			Content: "this is post 6",
			UserID:  1,
		},
	}

	var comments = []Comment{
		{
			Content: "1th comment of post 1",
			UserID:  1,
			PostID:  1,
		},
		{
			Content: "2th comment of post 1",
			UserID:  1,
			PostID:  1,
		},
	}

	fmt.Println(users)
	// if res := db.Debug().CreateInBatches(&users, 1); res.Error != nil {
	// 	fmt.Println("用户数据初始化失败", res.Error)
	// } else {
	// 	fmt.Println("用户数据初始化成功")
	// }

	if res := db.Debug().CreateInBatches(&posts, 1); res.Error != nil {
		fmt.Println("文章数据初始化失败：", res.Error)
	} else {
		fmt.Println("文章数据初始化成功")
	}

	if res := db.Debug().CreateInBatches(&comments, 1); res.Error != nil {
		fmt.Println("评论数据初始化失败：", res.Error)
	} else {
		fmt.Println("评论数据初始化成功")
	}

}

/*
*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。

*
*/
func getPostWithCommentByUserId(db *gorm.DB, userID uint) ([]Post, []Post) {
	// 使用外键查询
	var posts []Post
	if err := db.Debug().Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		fmt.Println("文章查询失败（外键关联评论）：", err)
	} else {
		fmt.Printf("查到的文章和评论%v：", posts)
	}

	// 手动关联查询评论
	var posts2 []Post
	if err := db.Debug().Where("user_id = ?", userID).Find(&posts2).Error; err != nil {
		fmt.Println("文章查询失败（手动关联）：", err)
	}
	// 手动添加评论
	for i := 0; i < len(posts2); i++ {
		if err := db.Debug().Where("post_id = ? and user_id = ?", posts2[i].ID, userID).Find(&posts2[i].Comments).Error; err == nil {
			fmt.Printf("查到的文章和评论(手动关联)%v：", posts2[i])
		}
	}

	return posts, posts2

}

/**

题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
**/

/**
钩子函数的签名是固定的, 必须返回 error 的原因:
1、错误处理：如果钩子函数执行石板，gorm 需要知道以便回滚事务
2、事务控制：返回错误会导致整个操作的回滚
3、一致性保证：确保钩子函数的执行结果影响主操作
**/

func (post *Post) AfterCreate(tx *gorm.DB) error {
	var count int64
	if err := tx.Debug().Model(&Post{}).Where("user_id = ?", post.UserID).Count(&count).Error; err != nil {
		fmt.Println("统计文章总数错误：", err)
		return err
	}

	if err := tx.Debug().Model(&User{}).Where("id = ?", post.UserID).Update("post_count", count).Error; err != nil {
		fmt.Println("更新用户文章数失败：", err)
		return err
	}
	return nil
}

// 删除评论
func deleteCommentById(db *gorm.DB, id uint) {

	// 方法1： 此方法删除评论，在删除的钩子方法中拿不到 Comment 的实例信息
	// if err := db.Debug().Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
	// 	fmt.Println("删除评论失败：", err)
	// }

	// 方法2 ： 此方法可以在 删除的钩子方法中拿到 Comment 的实现信息
	// 先查出来要删除的评论实体
	var delCondi Comment
	if err := db.Debug().Model(&Comment{}).Where("id = ?", id).First(&delCondi).Error; err != nil {
		fmt.Println("查询评论失败：", err)
		return
	}
	fmt.Println("即将删除的评论：", delCondi)
	// 传入实体指针，这样钩子函数才能拿到评论的完整信息，以便根据评论的文章id 更新文章评论数
	if err := db.Debug().Delete(&delCondi).Error; err != nil {
		fmt.Println("删除评论失败：", err)
	}
}

// 删除评论的钩子函数
// AfterDelete BeforeDelete 钩子中，都可以实现更新评论数量，前提是删除数传入的评论实例包含：文章id
func (comment *Comment) AfterDelete(tx *gorm.DB) error {

	fmt.Println("删除的评论：", *comment)
	var postCommentCount int64
	if err := tx.Debug().Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&postCommentCount).Error; err != nil {
		fmt.Println("统计文章评论数错误：", err)
		return err
	}

	postCommentCount = postCommentCount - 1
	if postCommentCount < 0 {
		postCommentCount = 0
	}

	if err := tx.Debug().Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_count", postCommentCount).Error; err != nil {
		fmt.Println("更新文章评论数错误：", err)
		return err
	}

	if postCommentCount == 0 {
		if err := tx.Debug().Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_status", "无评论").Error; err != nil {
			fmt.Println("更新文章评论状态失败：", err)
			return err
		}
	}

	return nil

}

/**

GO 中的钩子函数（Hook Functions） 是指在特定事件发生时调用的函数。
在gorom 中，钩子函数是模型生命周期方法，可以在数据库操作的特定阶段自动执行自定义的逻辑。

// 创建相关的钩子
func (m *Model) BeforeCreate(tx *gorm.DB) error {} // 创建前
func (m *Model) AfterCreate(tx *gorm.DB) error {}  // 创建后

// 更新相关的钩子
func (m *Model) BeforeUpdate(tx *gorm.DB) error {} // 更新前
func (m *Model) AfterUpdate(tx *gorm.DB) error {}  // 更新后

// 删除相关的钩子
func (m *Model) BeforeDelete(tx *gorm.DB) error {} // 删除前
func (m *Model) AfterDelete(tx *gorm.DB) error {}  // 删除后

// 查询相关的钩子
func (m *Model) AfterFind(tx *gorm.DB) error {}    // 查询后
**/
