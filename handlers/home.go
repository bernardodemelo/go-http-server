package handlers

import (
 "net/http"
 "fmt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Welcome to our custom GO HTTP server!")
}