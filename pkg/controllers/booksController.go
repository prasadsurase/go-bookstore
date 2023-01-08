package booksController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prasadsurase/go-bookstore/pkg/models"
	"github.com/prasadsurase/go-bookstore/pkg/utils"
)

func Books(w http.ResponseWriter, r *http.Request) {
	books := models.GetAll()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func Book(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book, err := models.GetById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		// w.Write()
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, _ := json.Marshal(book)
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book *models.Book
	utils.ParseBody(r, book)
	book, err := book.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		// w.Write()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(book)
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(r, book)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	existingBook, db, err := models.GetById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("error while finding book by id")
	}
	if existingBook == nil {
		fmt.Println("Unable to find book by id")
	}
	if book.Name != "" {
		existingBook.Name = book.Name
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Publication != "" {
		existingBook.Publication = book.Publication
	}
	db.Save(&existingBook)
	res, _ := json.Marshal(existingBook)
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book, err := models.DeleteById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("error while deleting book by id")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(book)
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
