package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexapps/cloud-native-go/api/util"
	"github.com/alexapps/cloud-native-go/model"
	bookService "github.com/alexapps/cloud-native-go/storage/moc"
)

/**
  Handlers used for REST needs
*/

type BookHandler struct {
	bs *bookService.BookService
}

func InitBookHandler() *BookHandler {
	return &BookHandler{
		bs: bookService.InitBookService(),
	}
}

// BooksHandleFunc processing requests "/api/books"
func (bh *BookHandler) BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books, _ := bh.bs.AllBooks()
		util.WriteJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := util.FromJSON(body)
		created, isbn, _ := bh.bs.Create(&book)
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
func (bh *BookHandler) CreateBook(book model.Book) (string, bool) {
	if _, ok := bh.bs.GetBooksMap()[book.ISBN]; ok {
		return "", false
	}
	bh.bs.SetBookMapValue(book.ISBN, &book)
	return book.ISBN, true
}

func (bh *BookHandler) BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	books, _ := bh.bs.AllBooks()
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
