package api

import (
	"github.com/TomBowyerResearchProject/common/verification"
	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", healthz)
	})

	// r.With(verification.VerifyToken()).Route("/notification_system", func(r chi.Router) {
	// 	r.Post("/", createNotification)
	// })

	r.With(verification.VerifyJTW()).Route("/notification", func(r chi.Router) {
		r.Post("/", createNotification)
		r.Get("/", getNotificationList)
	})

	return r
}
