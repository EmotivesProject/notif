package consumer

import (
	"encoding/json"
	"notif/internal/db"
	"notif/model"

	"github.com/TomBowyerResearchProject/common/logger"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var (
	topic         = "NOTIF"
	kafkaConsumer *kafka.Consumer
)

func Init() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "notifs",
	})
	if err != nil {
		logger.Error(err)
	}

	kafkaConsumer = consumer
}

func Run() {
	err := kafkaConsumer.Subscribe(topic, nil)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Connect to kafka")
	defer kafkaConsumer.Close()

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			createNotificationFromKafkaMessage(msg.Value)
		} else {
			logger.Error(err)
		}
	}
}

func createNotificationFromKafkaMessage(event []byte) {
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
