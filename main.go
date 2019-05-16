package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author info
// Interestingly, if the variable in a struct not started with
// uppercase letter, JSON package cannot see it! It's only exported
// when written LikeThis.
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func buildDB() {
	books = append(books, Book{
		ID:    "1",
		Title: "1984",
		Author: &Author{
			FirstName: "George",
			LastName:  "Orwell",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Title: "Animal Farm",
		Author: &Author{
			FirstName: "George",
			LastName:  "Orwell",
		},
	})
}

func main() {
	// router variable from mux
	router := mux.NewRouter()
	buildDB()
	// endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getOneBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("GET")
	// run server
	(http.ListenAndServe(":8000", router))
}
