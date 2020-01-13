package main

import (
	"app/api/books"
	"app/api/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	defer config.Producer.Stop()
	/*
		StrictSlash defines the slash behavior for new routes.
		When true, if the route path is "/path/", accessing "/path" will redirect to the former and vice versa.
	*/
	router.HandleFunc("/books", books.Index)
	router.HandleFunc("/books/create", books.Create).Methods("POST")
	router.HandleFunc("/books/{id}", books.ShowOne).Methods("GET")
	router.HandleFunc("/books/{id}", books.Update).Methods("PATCH")
	router.HandleFunc("/books/{id}", books.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
