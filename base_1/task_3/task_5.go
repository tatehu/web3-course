package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

// 1. 定义模型
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Posts []Post // 一对多关系：一个用户有多篇文章
}

type Post struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	UserID   uint      // 外键（属于哪个用户）
	Comments []Comment // 一对多关系：一篇文章有多个评论
}

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint // 外键（属于哪篇文章）
}

func main() {
	dsn := "root:root@tcp(192.168.159.168:33306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}

	// 自动迁移生成表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("自动建表失败:" + err.Error())
	}

	// 1. 初始化数据
	// 批量插入用户
	users := []User{
		{Name: "张三"},
		{Name: "李四"},
	}
	db.Create(&users)

	// 批量插入文章
	posts := []Post{
		{Title: "GORM介绍", Content: "GORM是Go的ORM框架", UserID: users[0].ID},
		{Title: "Go并发编程", Content: "介绍Go的goroutine和channel", UserID: users[0].ID},
		{Title: "SQL优化", Content: "索引优化与执行计划分析", UserID: users[1].ID},
	}
	db.Create(&posts)

	// 批量插入评论
	comments := []Comment{
		{Content: "很有用的文章！", PostID: posts[0].ID},
		{Content: "讲得很清楚", PostID: posts[0].ID},
		{Content: "学习了！", PostID: posts[1].ID},
		{Content: "索引讲解很细", PostID: posts[2].ID},
	}
	db.Create(&comments)

	// 2. 一对多查询
	// 查询某个用户的所有文章
	var userWithPosts User
	db.Preload("Posts").First(&userWithPosts, users[0].ID)
	fmt.Printf("用户: %s 的所有文章：\n", userWithPosts.Name)
	for _, p := range userWithPosts.Posts {
		fmt.Printf("  - %s（内容：%s）\n", p.Title, p.Content)
	}

	// 查询某篇文章及其所有评论
	var postWithComments Post
	db.Preload("Comments").First(&postWithComments, posts[0].ID)
	fmt.Printf("文章: %s 的所有评论：\n", postWithComments.Title)
	for _, c := range postWithComments.Comments {
		fmt.Printf("  - %s\n", c.Content)
	}

}
