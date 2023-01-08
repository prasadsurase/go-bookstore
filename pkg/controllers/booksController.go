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
	books, _ := models.GetAll()
	res, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Something went wrong. Please try again later")
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func Book(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book, err := models.GetById(id)
	if err != nil {
		fmt.Println("Something went wrong. Please try again later")
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		res, err := json.Marshal(book)
		if err != nil {
			fmt.Println("Something went wrong. Please try again later")
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	book, err := book.CreateBook()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Unable to create book. Please try again later.")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(book)
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book, err := models.GetById(id)
	if err != nil {
		fmt.Println("Something went wrong. Please try again later")
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		utils.ParseBody(r, book)
		err = book.Save()
		if err != nil {
			fmt.Println("Unable to save. Please try again later")
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			res, _ := json.Marshal(book)
			w.Write(res)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing id")
		return
	}
	book, err := models.DeleteById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error while deleting book by id")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(book)
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
