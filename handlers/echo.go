package handlers

import (
	"fmt"
	"net/http"
)

// Echo -
func (bh *BookHandler) Echo(w http.ResponseWriter, r *http.Request) {
	// Extarct the input message. The first one
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

// Echo -
func (bh *BookHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, it seems to be good")
}
