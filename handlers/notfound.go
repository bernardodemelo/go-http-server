package handlers

import (
 "net/http"
 "fmt"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Page not found")
   }