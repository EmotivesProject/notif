// +build integration

package api_test

import (
	"fmt"
	"net/http"
	"notif/test"
	"testing"

	"github.com/TomBowyerResearchProject/common/notification"
	commonTest "github.com/TomBowyerResearchProject/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRouterGetNotif(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	test.CreateNotification(t, username, token)

	req, _ := http.NewRequest("GET", test.TS.URL+"/notification", nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}

func TestRouterGetNotifTpe(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	test.CreateNotification(t, username, token)

	url := fmt.Sprintf("%s/notification/type/%s", test.TS.URL, notification.Like)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}
