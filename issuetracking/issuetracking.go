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
	AddIssueLabel(common.Issue, string) error
	RemoveIssueLabel(common.Issue, string) error
}

func GetClient() IssueTrackingClient {
	return GitHub{
		GitHubService: service.GitHub{},
	}
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
