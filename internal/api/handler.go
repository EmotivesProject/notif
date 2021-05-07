package api

import (
	"encoding/json"
	"net/http"
	"notif/internal/db"
	"notif/model"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getNotificationList(w http.ResponseWriter, r *http.Request) {
	page := findPage(r)

	username, ok := r.Context().Value(verification.UserID).(string)
	if !ok {
		response.MessageResponseJSON(w, http.StatusOK, response.Message{
			Message: "Failed to convert",
		})

		return
	}

	notifications := db.FindNotificationsByUsername(username, page)

	response.ResultResponseJSON(w, http.StatusOK, notifications)
}

func createNotification(w http.ResponseWriter, r *http.Request) {
	notification := &model.Notification{}
	if err := json.NewDecoder(r.Body).Decode(notification); err != nil {
		logger.Error(err)
		response.MessageResponseJSON(w, http.StatusBadRequest, response.Message{Message: err.Error()})

		return
	}

	notification.ID = primitive.NewObjectID()
	notification.CreatedAt = time.Now()
	notification.Seen = false

	if err := db.CreateNotification(notification); err != nil {
		logger.Error(err)
		response.MessageResponseJSON(w, http.StatusBadRequest, response.Message{Message: err.Error()})

		return
	}

	response.ResultResponseJSON(w, http.StatusOK, notification)
}
