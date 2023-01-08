package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prasadsurase/go-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	router.Handle("/", router)
	fmt.Println("Starting server at port 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", router))

}
