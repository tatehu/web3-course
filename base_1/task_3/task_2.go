package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务
*/

// Account 表结构
type Account struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Balance int
}

// Transaction 表结构
type Transaction struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	FromAccountID uint
	ToAccountID   uint
	Amount        int
}

func main() {
	dsn := "root:root@tcp(192.168.159.168:33306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移表结构
	db.AutoMigrate(&Account{}, &Transaction{})

	// 示例数据，假设A的id=1，B的id=2
	fromID := uint(1)
	toID := uint(2)
	amount := 100

	err = Transfer(db, fromID, toID, amount)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}

// 转账事务函数
func Transfer(db *gorm.DB, fromID, toID uint, amount int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var from Account
		// 查询转出账户 - 使用FOR UPDATE加锁
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&from, fromID).Error; err != nil {
			return err
		}
		// 检查余额是否足够
		if from.Balance < amount {
			return errors.New("账户余额不足")
		}

		// 扣除A账户余额
		if err := tx.Model(&Account{}).Where("id = ?", fromID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		// 增加B账户余额
		if err := tx.Model(&Account{}).Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}
		// 记录转账信息
		txRecord := Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(&txRecord).Error; err != nil {
			return err
		}
		return nil // 提交事务
	})
}
