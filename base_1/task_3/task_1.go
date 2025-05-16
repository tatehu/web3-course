package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

// Student 定义学生模型
type Student struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(50)"`
	Age   int    `gorm:"type:int"`
	Grade string `gorm:"type:varchar(50)"`
}

func main() {
	// 连接数据库
	dsn := "root:root@tcp(192.168.159.168:33306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&Student{})

	// 1. 插入记录
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	result := db.Create(&student)
	if result.Error != nil {
		fmt.Println("插入失败:", result.Error)
	} else {
		fmt.Println("插入成功")
	}

	// 2. 查询年龄大于18岁的学生
	var students []Student
	result = db.Where("age > ?", 18).Find(&students)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
	} else {
		fmt.Println("查询到的学生:", students)
	}

	// 3. 更新张三的年级
	result = db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	if result.Error != nil {
		fmt.Println("更新失败:", result.Error)
	} else {
		fmt.Printf("更新成功，影响的记录数: %d\n", result.RowsAffected)
	}

	// 4. 删除年龄小于15岁的学生
	result = db.Where("age < ?", 15).Delete(&Student{})
	if result.Error != nil {
		fmt.Println("删除失败:", result.Error)
	} else {
		fmt.Printf("删除成功，影响的记录数: %d\n", result.RowsAffected)
	}
}
