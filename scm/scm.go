package scm

import (
	"github.com/jtuttle/dev-flow/common"
)

type SourceControlManagement interface {
	GetPullRequest(string) *PullRequest
	CreatePullRequest(common.Issue) *PullRequest
	MergePullRequest(*PullRequest) bool
}

func (pr PullRequest) String() string {
	return pr.URL
}

func GetClient() SourceControlManagement {
	return GitHub { }
}

type PullRequest struct {
	Number int
	URL string
	Base string
	Mergeable bool
}
