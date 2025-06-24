package routes

import (
	"http-go/handlers"
	"http-go/server"
)

func RegisterRoutes(s *server.Server) {
	r := s.GetRouter()
	r.GET("/", handlers.HomeHandler)
	r.GET("/roller-coaster", handlers.TimeHandler)
	r.NotFound(handlers.NotFoundHandler)
}
