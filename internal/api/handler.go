package api

import (
	"net/http"
	"notif/internal/db"
	"notif/messages"

	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: messages.HealthResponse})
}

func getNotificationList(w http.ResponseWriter, r *http.Request) {
	page := findPage(r)
	username := r.Context().Value(verification.UserID).(string)

	notifications := db.FindNotificationsByUsername(username, page)

	response.ResultResponseJSON(w, http.StatusOK, notifications)
}
