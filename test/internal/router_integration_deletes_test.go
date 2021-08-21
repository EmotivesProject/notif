// +build integration

package api_test

import (
	"fmt"
	"net/http"
	"notif/test"
	"testing"

	commonTest "github.com/TomBowyerResearchProject/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRouterDeleteNotifPostID(t *testing.T) {
	test.SetUpIntegrationTest()

	username, token := commonTest.CreateNewUser(t, test.UaclUserEndpoint)

	test.CreateNotification(t, username, token)

	url := fmt.Sprintf("%s/notification/post/%d", test.TS.URL, 1)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", token)

	r, _, _ := commonTest.CompleteTestRequest(t, req)

	assert.EqualValues(t, r.StatusCode, http.StatusOK)

	test.TearDownIntegrationTest()
}
