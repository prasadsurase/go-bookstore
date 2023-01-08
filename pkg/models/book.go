package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/prasadsurase/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"NOT NULL"json:"name"`
	Author      string `gorm:"NOT NULL"json:"author"`
	Publication string `gorm:"NOT NULL"json:"publication"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(&b)
	if result.Error == nil {
		return b, nil
	} else {
		return nil, result.Error
	}
}

func GetAll() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error == nil {
		return books, nil
	} else {
		return nil, result.Error
	}
}

func GetById(id int64) (*Book, error) {
	var book *Book
	result := db.First(&book, id)
	if result.Error == nil {
		return book, nil
	} else {
		return nil, result.Error
	}
}

func DeleteById(id int64) (*Book, error) {
	var book *Book
	result := db.Where("id = ?", id).First(&book).Delete(book)
	if result.Error == nil {
		return book, nil
	} else {
		return nil, result.Error
	}
}

func (b *Book) Save() error {
	result := db.Save(b)
	return result.Error
}

func (b *Book) Delete() error {
	result := db.Delete(Book{}, b.ID)
	return result.Error
}
