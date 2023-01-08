package routes

import (
	"github.com/gorilla/mux"
	booksController "github.com/prasadsurase/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", booksController.Books).Methods("GET")
	router.HandleFunc("/book", booksController.Book).Methods("GET")
	router.HandleFunc("/books", booksController.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", booksController.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", booksController.DeleteBook).Methods("DELETE")
}
