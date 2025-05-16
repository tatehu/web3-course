package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
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

	// ------- 查询某个用户发布的所有文章及其对应的评论信息 -------
	var user User
	userID := 1 // 假设查id=1的用户，可根据实际情况调整
	err = db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Println("用户不存在或查询失败:", err)
	} else {
		fmt.Printf("用户[%s]的所有文章及评论：\n", user.Name)
		for _, post := range user.Posts {
			fmt.Printf("  文章：%s\n", post.Title)
			for _, c := range post.Comments {
				fmt.Printf("    评论：%s\n", c.Content)
			}
		}
	}

	// ------- 查询评论数量最多的文章信息 -------
	var mostCommented Post
	var commentCount int64

	// 先查出评论数最多的文章ID
	type Result struct {
		PostID uint
		Count  int64
	}
	var res Result
	db.Model(&Comment{}).
		Select("post_id, COUNT(*) as count").
		Group("post_id").
		Order("count DESC").
		Limit(1).
		Scan(&res)

	if res.PostID != 0 {
		db.Preload("Comments").First(&mostCommented, res.PostID)
		commentCount = res.Count
		fmt.Printf("\n评论最多的文章：%s（%d条评论）\n", mostCommented.Title, commentCount)
		for _, c := range mostCommented.Comments {
			fmt.Printf("  - %s\n", c.Content)
		}
	} else {
		fmt.Println("\n未找到评论数据。")
	}

}
