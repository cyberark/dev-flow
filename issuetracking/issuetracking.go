package issuetracking

import (
	"strings"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
)

type IssueTrackingClient interface {
	GetCurrentUserLogin() (string, error)
	GetUserRealName(string) (string, error)
	GetIssues() ([]common.Issue, error)
	GetIssue(string) (*common.Issue, error)
	AssignIssue(int, string) error
	AddIssueLabel(int, string) error
	RemoveIssueLabel(int, string) error
}

func GetClient(repo common.Repo) IssueTrackingClient {
	return GitHub{
		Repo: repo,
		GitHubService: service.GitHub{},
	}
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
