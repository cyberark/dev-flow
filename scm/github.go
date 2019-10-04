package scm

import (
	"fmt"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
)

type GitHub struct{
	Repo common.Repo
}

func newGitHubClient() service.GitHub {
	return service.GitHub{}
}

func (gh GitHub) GetPullRequest(branchName string) *PullRequest {
	ghPullRequests, err := newGitHubClient().GetPullRequests(gh.Repo, branchName)
	
	if err != nil {
		panic(err)
	}

	var pullRequestNum *int

	for _, ghPullRequest := range ghPullRequests {
		if *ghPullRequest.Head.Ref == branchName {
			pullRequestNum = ghPullRequest.Number
		}
	}

	var pullRequest *PullRequest

	if pullRequestNum != nil {
		// We call List and then Get because the Mergeable field is only
		// returned when retrieving single issues.
		ghPullRequest, err := newGitHubClient().GetPullRequest(gh.Repo, *pullRequestNum)
		
		if err != nil {
			panic(err)
		}

		pullRequest = gh.toCommonPullRequest(ghPullRequest)
	}

	return pullRequest
}

func (gh GitHub) CreatePullRequest(issue common.Issue, linkType string) *PullRequest {
	base := "master"
	head := issue.BranchName()
	title := issue.Title
	body := fmt.Sprintf("%s #%v", linkType, issue.Number)

	newPullRequest := &github.NewPullRequest{
		Base:  &base,
		Head:  &head,
		Title: &title,
		Body:  &body,
	}
	
	ghPullRequest, err := newGitHubClient().CreatePullRequest(gh.Repo, newPullRequest)
	
	if err != nil {
		panic(err)
	}

	err = newGitHubClient().AssignIssue(gh.Repo, *ghPullRequest.Number, issue.Assignee)

	if err != nil {
		panic(err)
	}

	return gh.toCommonPullRequest(ghPullRequest)
}

func (gh GitHub) AssignPullRequestReviewer(pr *PullRequest, reviewer string) {
	err := newGitHubClient().RequestReviewer(gh.Repo, pr.Number, reviewer)

	if err != nil {
		panic(err)
	}
}

func (gh GitHub) MergePullRequest(pr *PullRequest, mergeMethod string) bool {
	merged, err := newGitHubClient().MergePullRequest(gh.Repo, pr.Number, mergeMethod)

	if err != nil {
		panic(err)
	}

	return merged
}

func (gh GitHub) toCommonPullRequest(ghpr *github.PullRequest) *PullRequest {
	mergeable := false

	if ghpr.Mergeable != nil {
		mergeable = *ghpr.Mergeable
	}

	return &PullRequest{
		Number:    *ghpr.Number,
		Creator:   *ghpr.User.Login,
		Base:      *ghpr.Base.Ref,
		Mergeable: mergeable,
		URL:       *ghpr.HTMLURL,
	}
}
