package main

import (
	"log"
	"net/http"
	"os"

	"notif/internal/api"
	"notif/internal/db"

	"github.com/TomBowyerResearchProject/common/logger"
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

	db.ConnectDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	logger.Info("STARTING SERVER")

	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
