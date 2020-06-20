package main

import (
	"net/http"

	"github.com/alexapps/cloud-native-go/config"
	"github.com/alexapps/cloud-native-go/handlers"
)

func main() {

	// Init Configuration
	confInstance := config.InitConfiguration()
	// Init handlers
	bookHandler := handlers.InitBookHandler()
	// http://0.0.0.0:8084/
	http.HandleFunc("/", bookHandler.Index)
	// http://0.0.0.0:8084/api/echo?message=Cloud+Native+Go
	http.HandleFunc("/api/echo", bookHandler.Echo)
	http.HandleFunc("/api/books", bookHandler.BooksHandleFunc)
	http.HandleFunc("/api/books/", bookHandler.BookHandleFunc)
	http.ListenAndServe(confInstance.Port, nil)
}
