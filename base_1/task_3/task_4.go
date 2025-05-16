package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

// 1. 定义与 books 表对应的结构体
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 替换为你的数据库配置信息
	dsn := "root:root@tcp(192.168.159.168:33306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
	defer db.Close()

	// 2. 查询价格大于50元的书籍（可加复杂条件筛选）
	var expensiveBooks []Book
	query := `SELECT id, title, author, price FROM books WHERE price > ?`
	err = db.Select(&expensiveBooks, query, 50)
	if err != nil {
		log.Fatalln("查询书籍失败:", err)
	}

	// 3. 打印结果，类型安全
	fmt.Println("价格大于50元的书籍：")
	for _, book := range expensiveBooks {
		fmt.Printf("%+v\n", book)
	}
}
