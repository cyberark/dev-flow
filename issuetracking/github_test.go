package issuetracking

import (
	"errors"
	"testing"

	"github.com/cyberark/dev-flow/service"
)

func TestGetCurrentUser(t *testing.T) {
	client := GitHub{
		GitHubService: service.GitHubMock{},
	}

	userLogin, _ := client.GetCurrentUser()

	if userLogin != "octocat" {
		t.Errorf("expected %v, got: %v", "octocat", userLogin)
	}
}

func TestGetCurrentUserErr(t *testing.T) {
	service := service.GitHubMock{
		Error: errors.New("an error"),
	}
	
	client := GitHub{
		GitHubService: service,
	}

	_, err := client.GetCurrentUser()

	if err == nil {
		t.Errorf("expected error")
	}
}
