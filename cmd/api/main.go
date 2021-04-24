package main

import (
	"log"
	"net/http"
	"os"

	"notif/internal/api"
	"notif/internal/consumer"
	"notif/internal/db"

	commonKafka "github.com/TomBowyerResearchProject/common/kafka"
	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/middlewares"
	"github.com/TomBowyerResearchProject/common/verification"
	"github.com/joho/godotenv"
)

func main() {
	router := api.CreateRouter()

	logger.InitLogger("notif")

	verification.Init(verification.VerificationConfig{
		VerificationURL:     "http://uacl/authorize",
		AuthorizationSecret: "qutSecret",
	})

	middlewares.Init(middlewares.Config{
		AllowedOrigin:  "*",
		AllowedMethods: "GET,POST,OPTIONS",
		AllowedHeaders: "Accept, Content-Type, Content-Length, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header",
	})

	db.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	commonKafka.InitConsumer(commonKafka.ConfigConsumer{
		Topic:  "NOTIF",
		Server: "kafka:9092",
		Group:  "notif",
		Handle: consumer.CreateNotificationFromKafkaMessage,
	})
	go commonKafka.Run()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	logger.Info("STARTING SERVER")

	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
