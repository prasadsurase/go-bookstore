package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/prasadsurase/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate()
}

func (b *Book) CreateBook() (*Book, error) {
	// db.NewRecord(b)
	result := db.Create(&b)
	if result.Error == nil {
		return b, nil
	} else {
		// panic(result.Error)
		return nil, result.Error
	}
}

func GetAll() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetById(Id int64) (*Book, *gorm.DB, error) {
	var book Book
	result := db.Where("id = ?", Id).First(&book)
	if result.Error == nil {
		return &book, db, nil
	} else {
		// panic(result.Error)
		return nil, db, result.Error
	}
}

func DeleteById(Id int64) (*Book, error) {
	var book Book
	result := db.Where("id = ?", Id).First(&book).Delete(book)
	if result.Error == nil {
		return &book, nil
	} else {
		// panic(result.Error)
		return nil, result.Error
	}
}
