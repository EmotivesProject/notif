package api

import (
	"net/http"

	"github.com/TomBowyerResearchProject/common/middlewares"
	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
	"github.com/go-chi/chi"
)

func CreateRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", response.Healthz)
	})

	r.With(verification.VerifyToken()).Route("/internal_notification", func(r chi.Router) {
		r.Post("/", createNotification)
	})

	r.With(verification.VerifyJTW()).Route("/notification", func(r chi.Router) {
		r.Get("/", getNotificationList)
		r.Post("/", createNotification)

		r.Post("/{id}", updateNotificationToSeen)

		r.Delete("/post/{id}", removeNotificationsByPostID)

		r.Get("/type/{type_name}", getNotificationsByType)

		r.Post("/link/{link}/username/{username}", updateNotificationsToSeen)
	})

	return r
}
