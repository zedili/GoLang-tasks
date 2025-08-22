package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3-db?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, gormErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if gormErr != nil {
		panic(gormErr)

	}

	sqlxDb, sqlxErr := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/web3-db?charset=utf8mb4&parseTime=True&loc=Local")
	if sqlxErr != nil {
		panic(sqlxErr)
	}
	fmt.Println(gormDb, sqlxDb)

	// gorm
	// 基本curd操作
	// insert(db, 14)
	// insert(db, 20)
	// query(db)
	// update(db)
	// delete(db)

	// 事务语句
	// initAccount(db)
	// Transfer(db, 2, 1, 100)

	// sqlx
	// 使用 gorm 进行初始化
	// sqlxInitTable(gormDb)
	// sqlxQueryEmployee1(sqlxDb, "技术部")
	// sqlxQueryEmployee2(sqlxDb)
	queryBooksOfPriceGrandThan50Yuan(sqlxDb)

	// gorm 进阶
	// gormAdvanceInitialTable(gormDb)
	// getPostWithCommentByUserId(gormDb, 1)
	// deleteCommentById(gormDb, 9)

}
