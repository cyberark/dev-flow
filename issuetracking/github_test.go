package issuetracking_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/service/mocks"
	"github.com/cyberark/dev-flow/testutils"
	"github.com/cyberark/dev-flow/versioncontrol"
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
		"success": { result: "monalisa octocat", err: nil },
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

func TestGetIssues(t *testing.T) {
	tests := map[string]struct{
		result []common.Issue
		err error
	}{
		"success": {
			result: []common.Issue{
				{
					URL: "https://github.com/octocat/Hello-World/issues/1347",
					Number: 1347,
					Title: "Found a bug",
					Assignee: "octocat",
					Labels: []string{"bug"},
				},
			},
			err: nil,
		},
		"propagates error": { result: nil, err: errors.New("an error") },
	}

	for _, test := range tests {
		var ghIssuesPtrs []*github.Issue
		
		if test.result != nil {
			ghIssues := make([]github.Issue, 0)
			
			testutils.LoadFixture("testdata/github_issues.json", &ghIssues)
			
			for i := 0; i < len(ghIssues); i++ {
				ghIssuesPtrs = append(ghIssuesPtrs, &ghIssues[i])
			}
		}

		// TODO: set up a mock for the git helper so that we can stub
		// repo info instead of using cyberark/dev-flow in the tests.
		repo := versioncontrol.Git{}.Repo()
		
		mockService := &mocks.GitHubService{}
		mockService.On("GetIssues", repo).Return(ghIssuesPtrs, test.err)
		
		client := issuetracking.GitHub{
			GitHubService: mockService,
		}

		issues, err := client.GetIssues()
		
		assert.Equal(t, test.result, issues)
		assert.Equal(t, test.err, err)
	}
}

func TestGetIssue(t *testing.T) {
	tests := map[string]struct{
		result *common.Issue
		err error
	}{
		"success": {
			result: &common.Issue{
				URL: "https://github.com/octocat/Hello-World/issues/1347",
				Number: 1347,
				Title: "Found a bug",
				Assignee: "octocat",
				Labels: []string{"bug"},
			},
			err: nil,
		},
		"propagates error": { result: nil, err: errors.New("an error") },
	}

	for _, test := range tests {
		var ghIssue *github.Issue
		
		if test.result != nil {
			testutils.LoadFixture("testdata/github_issue.json", &ghIssue)
		}

		// TODO: set up a mock for the git helper so that we can stub
		// repo info instead of using cyberark/dev-flow in the tests.
		repo := versioncontrol.Git{}.Repo()
		issueNum := 1347
		
		mockService := &mocks.GitHubService{}
		mockService.On("GetIssue", repo, issueNum).Return(ghIssue, test.err)
		
		client := issuetracking.GitHub{
			GitHubService: mockService,
		}

		issue, err := client.GetIssue(strconv.Itoa(issueNum))
		
		assert.Equal(t, test.result, issue)
		assert.Equal(t, test.err, err)
	}
}

func TestAssignissue(t *testing.T) {
	tests := map[string]struct{
		err error
	}{
		"success": { err: nil },
		"propagates error": { err: errors.New("an error") },
	}

	for _, test := range tests {
		// TODO: set up a mock for the git helper so that we can stub
		// repo info instead of using cyberark/dev-flow in the tests.
		repo := versioncontrol.Git{}.Repo()
		issueNum := 123
		login := "octocat"

		mockService := &mocks.GitHubService{}
		mockService.On("AssignIssue", repo, issueNum, login).Return(test.err)

		client := issuetracking.GitHub{
			GitHubService: mockService,
		}
		
		err := client.AssignIssue(issueNum, login)
		
		assert.Equal(t, test.err, err)
	}
}

func TestAddIssueLabel(t *testing.T) {
	tests := map[string]struct{
		err error
	}{
		"success": { err: nil },
		"propagates error": { err: errors.New("an error") },
	}

	for _, test := range tests {
		repo := versioncontrol.Git{}.Repo()
		issueNum := 123
		labelName := "test"

		mockService := &mocks.GitHubService{}
		mockService.On("AddLabelToIssue", repo, issueNum, labelName).Return(test.err)

		client := issuetracking.GitHub{
			GitHubService: mockService,
		}
		
		err := client.AddIssueLabel(issueNum, labelName)
		
		assert.Equal(t, test.err, err)		
	}
}

func TestRemoveIssueLabel(t *testing.T) {
	tests := map[string]struct{
		err error
	}{
		"success": { err: nil },
		"propagates error": { err: errors.New("an error") },
	}

	for _, test := range tests {
		repo := versioncontrol.Git{}.Repo()
		issueNum := 123
		labelName := "test"

		mockService := &mocks.GitHubService{}
		mockService.On("RemoveLabelForIssue", repo, issueNum, labelName).Return(test.err)

		client := issuetracking.GitHub{
			GitHubService: mockService,
		}
		
		err := client.RemoveIssueLabel(issueNum, labelName)
		
		assert.Equal(t, test.err, err)		
	}
}
