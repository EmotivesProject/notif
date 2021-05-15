package test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"notif/internal/api"
	"notif/internal/db"
	"strings"
	"testing"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"github.com/TomBowyerResearchProject/common/notification"
	commonTest "github.com/TomBowyerResearchProject/common/test"
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

func CreateNotification(t *testing.T, username, token string) string {
	requestBody := strings.NewReader(
		fmt.Sprintf(
			`{\"username\": \"%s\",
			\"type\": \"%s\",
			\"title\": \"yo\",
			\"message\": \"messess\",
			\"link\":\"ye\",
			\"post_id\":1}`,
			username,
			notification.Like,
		),
	)

	req, _ := http.NewRequest("POST", TS.URL+"/notification", requestBody)
	req.Header.Add("Authorization", token)

	r, resultMap, _ := commonTest.CompleteTestRequest(t, req)
	r.Body.Close()

	return resultMap["id"].(string)
}
