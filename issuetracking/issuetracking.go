package issuetracking

import (
	"strings"

	"github.com/conjurinc/dev-flow/common"
)

type IssueTrackingClient interface {
	GetCurrentUser() string
	Issues() []common.Issue
	Issue(string) common.Issue
	AssignIssue(common.Issue, string)
	LabelIssue(common.Issue, string) error
}

func GetClient() IssueTrackingClient {
	return GitHub{}
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
