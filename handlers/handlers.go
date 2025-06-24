package handlers

import (
 "encoding/json"
 "fmt"
 "net/http"
 "time"
 "github.com/go-resty/resty/v2"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/plain")
 fmt.Fprintln(w, "Welcome to our custom HTTP server!")
}

func RollerCoasterHandler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")
 
 client := resty.New()
 resp,err := client.R()
	setHeader("Accept", "application/json")
	Get()
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/plain")
 w.WriteHeader(http.StatusNotFound)
 fmt.Fprintln(w, "404 - Page not found")
}