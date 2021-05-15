// +build integration

package api_test

import (
	"fmt"
	"net/http"
	"notif/test"
	"strings"
	"testing"

	"github.com/TomBowyerResearchProject/common/notification"
	commonTest "github.com/TomBowyerResearchProject/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRouterCreateNotif(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	requestBody := strings.NewReader(
		fmt.Sprintf(
			"{\"username\": \"%s\", \"type\": \"%s\", \"title\": \"yo\", \"message\": \"messess\", \"link\":\"ye\"}",
			username,
			notification.Like,
		),
	)

	req, _ := http.NewRequest("POST", test.TS.URL+"/notification", requestBody)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusCreated)

	test.TearDownIntegrationTest()
}

func TestRouterUpdateNotif(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	id := test.CreateNotification(t, username, token)

	url := fmt.Sprintf("%s/notification/%s", test.TS.URL, id)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}

func TestRouterUpdateNotifLinkUsername(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	test.CreateNotification(t, username, token)

	url := fmt.Sprintf(
		"%s/notification/link/ye/username/%s",
		test.TS.URL,
		username,
	)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}
