package main

import (
	"log"
	"net/http"
	"notif/internal/api"
	"notif/internal/db"
	"os"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/middlewares"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"github.com/TomBowyerResearchProject/common/verification"
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
		// nolint:lll
		AllowedHeaders: "Accept, Content-Type, Content-Length, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header",
	})

	err := commonMongo.Connect(commonMongo.Config{
		URI:    "mongodb://admin:admin@mongo_db:27017",
		DBName: db.DBName,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Info("STARTING SERVER")

	log.Fatal(http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), router))
}
