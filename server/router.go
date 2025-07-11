package server

import (
	"http-go/handlers"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/roller-coaster", func(r chi.Router) {
		r.Get("/", handlers.ListRollerCoasters)
		r.Post("/", handlers.CreateRollerCoaster)

		r.Route("/{rollerCoasterId}", func(r chi.Router) {
			r.Get("/", handlers.GetRollerCoaster)
		})
	})

	return r
}
