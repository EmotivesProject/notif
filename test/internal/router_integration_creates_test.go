// +build integration

package api_test

import (
	"notif/test"
	"testing"

	commonTest "github.com/TomBowyerResearchProject/common/test"
)

func TestRouterCreatePost(t *testing.T) {
	test.SetUpIntegrationTest()

	commonTest.CreateNewUser(t, "http://0.0.0.0:8082/user")

	test.TearDownIntegrationTest()
}
