package routes

import (
	"http-go/handlers"
	"http-go/server"
)

func RegisterRoutes() *server.Router {
	r := server.NewRouter()
	r.GET("/", handlers.HomeHandler)
	r.GET("/roller-coaster", handlers.RollerCoasterHandler)
	r.NotFound(handlers.NotFoundHandler)

	return r
}
