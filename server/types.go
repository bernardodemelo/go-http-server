package server

import (
	"net/http"
)

type Server struct {
	port   int
	router http.Handler
	server *http.Server
}
