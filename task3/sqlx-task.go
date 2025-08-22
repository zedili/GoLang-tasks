package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

/**

题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

**/

type Employees struct {
	// 如果没有指定字段类型，gorm 会尝试推断
	// 对于 string 类型，不同数据库的默认处理方式不同
	// mysql 中如果没有明确指定 varchar 长度，可能会生成无效的 sql。导致 gorm 表结构自动迁移失败
	// gorm 只能识别和导出 public 的字段（go 中首字母大写的字段为 public），private 字段会被 gorm 忽略
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(50);not null"`
	Department string `gorm:"type:varchar(100);not null"`
	Salary     int    `gorm:"not null"`
}

type Book struct {
	Id     uint `gorm:"primaryKey;autoIncrement"`
	Title  string
	Author string
	Price  float64
}

// 使用gorm 自动迁移表结构，初始化数据
func sqlxInitTable(db *gorm.DB) {
	// 自动迁移表格结构
	if err := db.Debug().AutoMigrate(&Employees{}); err != nil {
		panic(err)
	}

	if err := db.Debug().AutoMigrate(&Book{}); err != nil {
		panic(err)
	}

	// 插入初始化数据
	employees := []Employees{
		{
			Name:       "张三",
			Department: "技术部",
			Salary:     30000,
		},
		{
			Name:       "李四",
			Department: "业务部",
			Salary:     8000,
		},
		{
			Name:       "王五",
			Department: "技术部",
			Salary:     200000,
		},
	}

	books := []Book{
		{
			Title:  "张三的书",
			Author: "张三",
			Price:  69.8,
		},
		{
			Title:  "李四的书",
			Author: "李四",
			Price:  78.5,
		},
		{
			Title:  "王五的书",
			Author: "王五",
			Price:  49.5,
		},
	}

	if res := db.Debug().CreateInBatches(employees, 2); res.Error != nil {
		fmt.Println("employee 表数据初始化失败：", res.Error)
	}

	if res := db.Debug().CreateInBatches(books, 2); res.Error != nil {
		fmt.Println("book 表数据初始化失败：", res.Error)
	}

}

// 查询技术部的所有员工
func sqlxQueryEmployee1(db *sqlx.DB, depa string) {

	var result []Employees

	if err := db.Select(&result, "select id ,name,department,salary from employees where department = ?", depa); err != nil {
		fmt.Println("查询失败：", err)
	} else {
		// printf： 格式输出
		fmt.Printf("%s所有员工：%v\n", depa, result)
	}

}

// 查询工资最高的员工
func sqlxQueryEmployee2(db *sqlx.DB) {
	var result []Employees
	if err := db.Select(&result, "select id, name, department, salary from employees order by salary desc limit 1 "); err != nil {
		fmt.Println("查询失败：", err)
	} else {
		fmt.Println("工资最高的员工：", result)
	}
}

/**

题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

**/

func queryBooksOfPriceGrandThan50Yuan(db *sqlx.DB) {
	var books []Book
	if err := db.Select(&books, "select id, title, author, price from books where price > ?", 50); err != nil {
		fmt.Println("查询价格大于50元的书籍失败：", err)
	} else {
		fmt.Println("价格大于的50元的书籍：", books)
	}
}
