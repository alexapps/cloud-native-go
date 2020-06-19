package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alexapps/cloud-native-go/api"
)

func main() {
	// http://0.0.0.0:8084/
	http.HandleFunc("/", index)
	// http://0.0.0.0:8084/api/echo?message=Cloud+Native+Go
	http.HandleFunc("/api/echo", api.Echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

// Good practice to microservices to do not hardcore the configurable values
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8083"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// Body
	fmt.Fprintf(w, "Hello Go native Go")
}
