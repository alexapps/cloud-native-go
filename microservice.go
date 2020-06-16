package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// http://0.0.0.0:8084/
	http.HandleFunc("/", index)
	// http://0.0.0.0:8084/api/echo?message=Cloud+Native+Go
	http.HandleFunc("/api/echo", echo)
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

//
func echo(w http.ResponseWriter, r *http.Request) {
	// Extarct the input message. The first one
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// Body
	fmt.Fprintf(w, "Hello Go native Go")
}
