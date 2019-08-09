package issuetracking

import (
	"testing"

	"github.com/cyberark/dev-flow/service"
)

func TestGetCurrentUser(t *testing.T) {
	client := GitHub{
		GitHubClient: service.GitHubMock{}.GetClient(),
	}

	userLogin := client.GetCurrentUser()

	if userLogin != "octocat" {
		t.Errorf("Expected: %v, got: %v", "octocat", userLogin)
	}
}
