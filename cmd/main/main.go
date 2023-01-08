package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prasadsurase/go-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	router.Handle("/", router)
	http.ListenAndServe("localhost:9000", router)
}
