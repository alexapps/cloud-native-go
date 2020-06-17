package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
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
var Books = []Book{
	Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
	Book{Title: "Somthing Go", Author: "JJ Booo", ISBN: "4433444"},
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
