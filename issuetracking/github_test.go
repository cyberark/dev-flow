package issuetracking

import (
	"testing"

	"github.com/cyberark/dev-flow/testutils"
	"github.com/jarcoal/httpmock"
)

func TestGetCurrentUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	response := testutils.LoadJsonFile("testdata/github_user.json")

	httpmock.RegisterResponder("GET", "https://api.github.com/user",
		httpmock.NewStringResponder(200, response))

	client := GitHub{}

	userLogin := client.GetCurrentUser()

	if userLogin != "octocat" {
		t.Errorf("Expected: %v, got: %v", "octocat", userLogin)
	}
}
