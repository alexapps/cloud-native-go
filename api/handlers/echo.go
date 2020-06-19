package handlers

import (
	"fmt"
	"net/http"
)

//
func (bs *BookService) Echo(w http.ResponseWriter, r *http.Request) {
	// Extarct the input message. The first one
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
