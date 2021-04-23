package api

import (
	"github.com/TomBowyerResearchProject/common/middlewares"
	"github.com/TomBowyerResearchProject/common/verification"
	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", healthz)
	})

	r.With(verification.VerifyJTW()).Route("/notification", func(r chi.Router) {
		r.Get("/", getNotificationList)
	})

	return r
}
