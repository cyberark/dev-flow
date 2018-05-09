package scm

import (
	"github.com/conjurinc/dev-flow/common"
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
	Creator string
	Base string
	Mergeable bool
	URL string
}
