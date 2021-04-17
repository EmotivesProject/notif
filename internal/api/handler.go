package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notif/internal/db"
	"notif/model"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: "Health ok"})
}

func createNotification(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(verification.UserID)
	var notification model.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		logger.Error(err)
		response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: err.Error()})
		return
	}

	notification.Username = fmt.Sprintf("%v", username)
	notification.Seen = false

	err = db.CreateNotification(&notification)
	if err != nil {
		logger.Error(err)
		response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: err.Error()})
		return
	}

	logger.Info("Created event")
	response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: "All good"})
}

func getNotificationList(w http.ResponseWriter, r *http.Request) {
	skip := findSkip(r)
	username := r.Context().Value(verification.UserID)

	notifications := db.FindNotificationsByUsername(fmt.Sprintf("%v", username), skip)

	response.ResultResponseJSON(w, http.StatusOK, notifications)
}
