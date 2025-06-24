package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	port   int
	router http.Handler
	server *http.Server
}

func NewServer(port int, router http.Handler) *Server {
	return &Server{
		port:   port,
		router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	fmt.Printf("Server starting on port %d\n", s.port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
