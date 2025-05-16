package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

// 模型定义
type User struct {
	ID        uint
	Name      string
	PostCount int // 统计该用户的文章数量
	Posts     []Post
}

type Post struct {
	ID            uint
	Title         string
	Content       string
	UserID        uint
	CommentStatus string // 评论状态（如"有评论"、"无评论"）
	Comments      []Comment
}

// -------- Post 钩子函数 --------
// 文章创建后，自动增加用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

type Comment struct {
	ID      uint
	Content string
	PostID  uint
}

// -------- Comment 钩子函数 --------
// 评论删除后，如果该文章评论数为0，修改文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

func main() {
	dsn := "root:root@tcp(192.168.159.168:33306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	// 自动建表
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic("自动建表失败: " + err.Error())
	}

	//创建用户、发文章、发/删评论
	user := User{Name: "王五"}
	db.Create(&user)

	post := Post{Title: "GORM钩子函数", Content: "讲解AfterCreate/AfterDelete", UserID: user.ID, CommentStatus: "无评论"}
	db.Create(&post)
	fmt.Println("创建后用户文章数量：")
	db.First(&user, user.ID)
	fmt.Println(user.PostCount) // 应为1

	comment := Comment{Content: "学习到了！", PostID: post.ID}
	db.Create(&comment)
	db.Model(&Post{}).Where("id = ?", post.ID).Update("comment_status", "有评论")

	// 删除评论验证 AfterDelete 钩子
	db.Delete(&comment)
	db.First(&post, post.ID)
	fmt.Println("删除评论后文章评论状态：", post.CommentStatus) // 应为"无评论"
}
