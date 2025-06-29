package server

import (
	"context"
	"fmt"
	"net/http"
)

func NewServer(port int, router http.Handler) *Server {
	addr := fmt.Sprintf(":%d", port)
	return &Server{
		port:   port,
		router: router,
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	fmt.Printf("Server starting on port %d\n", s.port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
