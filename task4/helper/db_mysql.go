package helper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDb *gorm.DB

func GerMyqlDb() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3-db?charset=utf8mb4&parseTime=True&loc=Local"
	if mysqlDb == nil {
		db, gormErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		mysqlDb = db

		if gormErr != nil {
			panic(gormErr)
		}
	}
	return mysqlDb.Debug()

}
