package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
func (bs *BookHandler) BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
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
		isbn, created := bs.bs.CreateBook(book)
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
func (bs *BookHandler) CreateBook(book model.Book) (string, bool) {
	if _, ok := Books[book.ISBN]; ok {
		return "", false
	}
	Books[book.ISBN] = book
	return book.ISBN, true
}
