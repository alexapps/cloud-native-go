package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// MOC Books
var Books = map[string]Book{
	"0123456789": Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", Description: "bla bla", ISBN: "0123456789"},
	"4433444":    Book{Title: "Somthing Go", Author: "JJ Booo", Description: "Some interesting book", ISBN: "4433444"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		fmt.Println("created ", isbn, created)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	if _, ok := Books[book.ISBN]; ok {
		return "", false
	}
	Books[book.ISBN] = book
	return book.ISBN, true
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	booksSlice := make([]Book, len(Books))
	index := 0
	for _, v := range Books {
		booksSlice[index] = v
		index++
	}
	return booksSlice
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
