package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Route")
}

type Book struct {
	Id     int64  `json:"id"`
	Author string `json:"author"`
	Name   string `json:"name"`
}

var books []Book

func allBooksRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "All Books Route ")
}

func insertBookRoute(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(body, &book)
	books = append(books, book)
	fmt.Fprint(w, "Insert Book Route "+book.Name)
	fmt.Fprint(w, books)
}

func main() {
	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/allBooks", allBooksRoute)
	http.HandleFunc("/insertBook", insertBookRoute)
	http.HandleFunc("/routes", handleRoute)

	http.ListenAndServe(":8080", nil)
}
