package common

import (
	"gin-demo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	dsn := "root:@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return DB
}

func GetDb() *gorm.DB {
	return DB
}
