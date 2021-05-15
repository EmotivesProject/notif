package test

import (
	"log"
	"math/rand"
	"net/http/httptest"
	"notif/internal/api"
	"notif/internal/db"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"github.com/TomBowyerResearchProject/common/verification"
)

var TS *httptest.Server

func SetUpIntegrationTest() {
	rand.Seed(time.Now().Unix())

	logger.InitLogger("notif")

	verification.Init(verification.VerificationConfig{
		VerificationURL: "http://0.0.0.0:8082/authorize",
	})

	err := commonMongo.Connect(commonMongo.Config{
		URI:    "mongodb://admin:admin@0.0.0.0:27015",
		DBName: db.DBName,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	router := api.CreateRouter()

	TS = httptest.NewServer(router)
}

func TearDownIntegrationTest() {
	TS.Close()
}
