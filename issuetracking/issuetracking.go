package issuetracking

import (
	"strings"

	"github.com/jtuttle/dev-flow/common"
)

type IssueTracking interface {
	GetCurrentUser() string
	Issues() []common.Issue
	Issue(string) common.Issue
	AssignIssue(common.Issue, string)
}

func GetClient() IssueTracking {
	return GitHub { }
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
