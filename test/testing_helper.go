package test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"notif/internal/api"
	"notif/internal/db"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/middlewares"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"github.com/TomBowyerResearchProject/common/notification"
	commonTest "github.com/TomBowyerResearchProject/common/test"
	"github.com/TomBowyerResearchProject/common/verification"
)

const (
	uaclEndpoint     = "http://0.0.0.0:8082"
	UaclUserEndpoint = uaclEndpoint + "/user"
)

var TS *httptest.Server

func SetUpIntegrationTest() {
	rand.Seed(time.Now().Unix())

	logger.InitLogger("notif", logger.EmailConfig{
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Level:    os.Getenv("EMAIL_LEVEL"),
	})

	verification.Init(verification.VerificationConfig{
		VerificationURL:     uaclEndpoint + "/authorize",
		AuthorizationSecret: "secret",
	})

	middlewares.Init(middlewares.Config{
		AllowedOrigins: "*",
		AllowedMethods: "GET,POST,OPTIONS,DELETE",
		AllowedHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token",
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

func CreateNotification(t *testing.T, username, token string) string {
	body := fmt.Sprintf(
		"{\"username\": \"%s\", \"type\": \"%s\", \"title\": \"yo\","+
			"\"message\": \"messess\", \"link\":\"ye\", \"post_id\":1}",
		username,
		notification.Like,
	)
	requestBody := strings.NewReader(body)

	req, _ := http.NewRequest("POST", TS.URL+"/internal_notification", requestBody)
	req.Header.Add("Authorization", "secret")

	r, resultMap, _ := commonTest.CompleteTestRequest(t, req)
	r.Body.Close()

	return resultMap["id"].(string)
}
