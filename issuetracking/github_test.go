package issuetracking_test

import (
	"errors"
	"testing"
	
	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/service/mocks"
	"github.com/cyberark/dev-flow/testutils"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUserLogin(t *testing.T) {
	tests := map[string]struct{
		result string
		err error
	}{
		"success":          { result: "octocat", err: nil },
		"propagates error": { result: "", err: errors.New("an error") },
	}

	for _, test := range tests {
		var user *github.User = nil

		if test.result != "" {
			user = &github.User{}
			testutils.LoadFixture("testdata/github_user.json", user)
		}

		mockService := &mocks.GitHubService{}
		mockService.On("GetUser", "").Return(user, test.err)

 		client := issuetracking.GitHub{
			GitHubService: mockService,
		}

		login, err := client.GetCurrentUserLogin()

		assert.Equal(t, test.result, login)
		assert.Equal(t, test.err, err)
	}
}

func TestGetUserRealName(t *testing.T) {
	tests := map[string]struct{
		result string
		err error
	}{
		"success":          { result: "monalisa octocat", err: nil },
		"propagates error": { result: "", err: errors.New("an error") },
	}

	for _, test := range tests {
		var user *github.User = nil

		if test.result != "" {
			user = &github.User{}
			testutils.LoadFixture("testdata/github_user.json", user)
		}

		mockService := &mocks.GitHubService{}
		mockService.On("GetUser", "octocat").Return(user, test.err)

 		client := issuetracking.GitHub{
			GitHubService: mockService,
		}

		login, err := client.GetUserRealName("octocat")

		assert.Equal(t, test.result, login)
		assert.Equal(t, test.err, err)
	}
}
