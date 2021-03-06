package scm

import (
	"github.com/cyberark/dev-flow/common"
)

type SourceControlManagementClient interface {
	GetPullRequest(string) *PullRequest
	CreatePullRequest(common.Issue, string) *PullRequest
	AssignPullRequestReviewer(*PullRequest, string)
	MergePullRequest(*PullRequest, string) bool
}

func (pr PullRequest) String() string {
	return pr.URL
}

func GetClient(repo common.Repo) SourceControlManagementClient {
	return GitHub{
		Repo: repo,
	}
}

type PullRequest struct {
	Number    int
	Creator   string
	Base      string
	Mergeable bool
	URL       string
}
