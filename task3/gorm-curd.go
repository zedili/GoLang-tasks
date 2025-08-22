package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/**
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
**/

type Students struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   uint
	Grade string
}

// 创建或者更新表结构
// 插入数据
func insert(db *gorm.DB, age uint) {
	// AutoMigrate ： 自动迁移功能，根据结构体创建或者更新数据库表
	//				  如果表不存在，会自动创建新表
	//                如果表存在但结构不匹配，会修改表结构以匹配
	err := db.AutoMigrate(&Students{})
	if err != nil {
		panic(err)
	}
	// 插入新纪录
	studen := Students{Name: "张三", Age: age, Grade: "三年级"}
	// Debug ：启用调试模式，会在控制台输出执行的 SQL 语句
	res := db.Debug().Create(&studen)
	if res.Error != nil {
		fmt.Println("插入失败：", res.Error)
	} else {
		fmt.Println("插入成功，", studen)
	}
}

// 查询年龄大于 18 的学生
func query(db *gorm.DB) {
	var resulets []Students
	res := db.Where("Age > ?", 18).Find(&resulets)
	if res.Error != nil {
		fmt.Println("查询失败：", res.Error)
	} else {
		fmt.Println("年龄大于18的学生:", resulets)
	}
}

// 更新张三的年级
func update(db *gorm.DB) {
	res := db.Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	if res.Error != nil {
		fmt.Println("更新失败：", res.Error)
	} else {
		fmt.Println("更新成功", res.RowsAffected)
	}
}

// 删除年龄小于 15 岁的学生
func delete(db *gorm.DB) {
	res := db.Where("age < ?", 15).Delete(&Students{})
	if res.Error != nil {
		fmt.Println("删除失败", res.Error)
	} else {
		fmt.Println("删除成功", res.RowsAffected)
	}
}

/**
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
**/

// 账户表
type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"not null"`
}

// 交易记录表
type Transation struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}

// 初始化账号
func initAccount(db *gorm.DB) {
	// 自动迁移创建表格
	err := db.AutoMigrate(&Account{}, &Transation{})
	if err != nil {
		panic(err)
	}

	// 初始化账号、余额
	accounts := []Account{
		{
			ID:      1,
			Balance: 100,
		},
		{
			ID:      2,
			Balance: 200,
		},
	}
	res := db.Debug().CreateInBatches(accounts, 2)
	if res.Error != nil {
		fmt.Println("账号初始化失败：", res.Error)
	} else {
		fmt.Println("账号初始化成功", accounts)
	}
}

func Transfer(db *gorm.DB, fromAccoutID uint, toAccountID uint, amount float64) {
	// 开启事务
	db.Debug().Transaction(func(tx *gorm.DB) error {
		// 1、检查转出账户余额是否充足
		var fromAccount Account
		if err := tx.First(&fromAccount, fromAccoutID).Error; err != nil {
			return err
		}

		// 2、判断余额是否充足
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}

		// 3、更新转出账户余额
		err := tx.Model(&Account{}).Where("id = ?", fromAccoutID).Update("balance", gorm.Expr("balance - ?", amount)).Error
		if err != nil {
			return err
		}

		// 4、更新转入账户余额
		if err := tx.Model(&Account{}).Where("id = ?", toAccountID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 5、 增加交易记录
		// 		ID uint `gorm:"primaryKey;autoIncrement"`
		// FromAccountID uint `gorm:"not null"`
		// ToAccountID uint `gorm:"not null"`
		// Amount float64 `gorm:"not null"`
		transation := Transation{
			FromAccountID: fromAccoutID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transation).Error; err != nil {
			fmt.Println("交易失败：", err)
			return err
		}
		return nil

	})
}
