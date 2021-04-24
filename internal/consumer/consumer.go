package consumer

import (
	"encoding/json"
	"notif/internal/db"
	"notif/model"

	"github.com/TomBowyerResearchProject/common/logger"
)

func CreateNotificationFromKafkaMessage(event []byte) {
	logger.Info("Got event saving it to the db")
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(event, &jsonMap)
	if err != nil {
		logger.Error(err)
		return
	}

	notification := model.CreateNotificationFromMap(jsonMap)

	err = db.CreateNotification(&notification)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Saved event")
}
