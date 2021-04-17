package api

import (
	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", healthz)
	})

	return r
}
