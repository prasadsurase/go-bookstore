package config

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/go-bookstore?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
