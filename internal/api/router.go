package api

import (
	"github.com/TomBowyerResearchProject/common/middlewares"
	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", response.Healthz)
	})

	r.With(verification.VerifyJTW()).Route("/notification", func(r chi.Router) {
		r.Get("/", getNotificationList)
		r.Post("/", createNotification)

		r.Post("/{link}/{username}", updateNotificationToSeen)
	})

	return r
}
