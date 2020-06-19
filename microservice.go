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
	booHandler := handlers.InitBookHandler()
	// http://0.0.0.0:8084/
	http.HandleFunc("/", index)
	// http://0.0.0.0:8084/api/echo?message=Cloud+Native+Go
	http.HandleFunc("/api/echo", booHandler.Echo)
	http.HandleFunc("/api/books", booHandler.BooksHandleFunc)
	http.HandleFunc("/api/books/", booHandler.BookHandleFunc)
	http.ListenAndServe(confInstance.Port, nil)
}
