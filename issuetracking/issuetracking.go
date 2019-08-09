package issuetracking

import (
	"strings"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
)

type IssueTrackingClient interface {
	GetCurrentUser() string
	GetUserRealName(string) string
	Issues() []common.Issue
	Issue(string) (common.Issue, error)
	AssignIssue(common.Issue, string)
	AddIssueLabel(common.Issue, string) error
	RemoveIssueLabel(common.Issue, string) error
}

func GetClient() IssueTrackingClient {
	return GitHub{
		GitHubClient: service.GitHub{}.GetClient(),
	}
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
