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

	login, _ := client.GetCurrentUserLogin()

	if login != "octocat" {
		t.Errorf("expected %v, got: %v", "octocat", login)
	}
}

func TestGetCurrentUserErr(t *testing.T) {
	service := service.GitHubMock{
		Error: errors.New("an error"),
	}
	
	client := GitHub{
		GitHubService: service,
	}

	login, err := client.GetCurrentUserLogin()

	if login != "" {
		t.Errorf("expected %v, got: %v", "", login)
	}
	
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGetUserRealName(t *testing.T) {
	client := GitHub{
		GitHubService: service.GitHubMock{},
	}

	realName, _ := client.GetUserRealName("octocat")

	if realName != "monalisa octocat" {
		t.Errorf("expected %v, got: %v", "octocat", realName)
	}
}

func TestGetUserRealNameErr(t *testing.T) {
	service := service.GitHubMock{
		Error: errors.New("an error"),
	}
	
	client := GitHub{
		GitHubService: service,
	}

	realName, err := client.GetUserRealName("octocat")

	if realName != "" {
		t.Errorf("expected %v, got: %v", "", realName)
	}
	
	if err == nil {
		t.Errorf("expected error")
	}
}
