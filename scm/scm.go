package scm

import (
	"github.com/cyberark/dev-flow/common"
)

type SourceControlManagementClient interface {
	GetPullRequest(string) *PullRequest
	CreatePullRequest(common.Issue) *PullRequest
	AssignPullRequestReviewer(*PullRequest, string)
	MergePullRequest(*PullRequest) bool
}

func (pr PullRequest) String() string {
	return pr.URL
}

func GetClient() SourceControlManagementClient {
	return GitHub{}
}

type PullRequest struct {
	Number    int
	Creator   string
	Base      string
	Mergeable bool
	URL       string
}
