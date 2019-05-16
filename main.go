package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// router variable from mux
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/api/books", getQuestions).Methods("GET")
	router.HandleFunc("/api/books/{id}", getCount).Methods("GET")
	router.HandleFunc("/api/books", getQuestions).Methods("GET")
	router.HandleFunc("/api/books/{id}", getCount).Methods("GET")
	router.HandleFunc("/api/books", getQuestions).Methods("GET")
	// run server
	(http.ListenAndServe(":8000", router))
}
