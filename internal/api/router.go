package api

import (
	"net/http"

	"github.com/EmotivesProject/common/middlewares"
	"github.com/EmotivesProject/common/response"
	"github.com/EmotivesProject/common/verification"
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

		r.Delete("/post/{id}", removeNotificationsByPostID)

		r.Delete("/like/post/{id}/user/{username}", removeLikeNotificationForUser)
	})

	r.With(verification.VerifyJTW()).Route("/notification", func(r chi.Router) {
		r.Get("/", getNotificationList)

		r.Post("/{id}", updateNotificationToSeen)

		r.Post("/link/username/{username}", updateNotificationsToSeen)
	})

	return r
}
