// +build integration

package api_test

import (
	"net/http"
	"notif/test"
	"testing"

	commonTest "github.com/EmotivesProject/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRouterGetNotif(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, test.UaclUserEndpoint)

	test.CreateNotification(t, username, token)

	req, _ := http.NewRequest("GET", test.TS.URL+"/notification", nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}
