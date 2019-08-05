package scm

import (
	"fmt"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/services"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHub struct{}

func newGitHubClient() services.GitHub {
	return services.GitHub{}.GetClient()
}

func (gh GitHub) GetPullRequest(branchName string) *PullRequest {
	repo := versioncontrol.GetClient().Repo()

	ghPullRequests, err := newGitHubClient().GetPullRequests(repo, branchName)
	
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
		ghPullRequest, err := newGitHubClient().GetPullRequest(repo, *pullRequestNum)
		
		if err != nil {
			panic(err)
		}

		pullRequest = gh.toCommonPullRequest(ghPullRequest)
	}

	return pullRequest
}

func (gh GitHub) CreatePullRequest(issue common.Issue, linkType string) *PullRequest {
	repo := versioncontrol.GetClient().Repo()

	base := "master"
	head := issue.BranchName()
	title := issue.Title
	body := fmt.Sprintf("%s #%v", linkType, *issue.Number)

	newPullRequest := &github.NewPullRequest{
		Base:  &base,
		Head:  &head,
		Title: title,
		Body:  &body,
	}
	
	ghPullRequest, err := newGitHubClient().CreatePullRequest(repo, newPullRequest)
	
	if err != nil {
		panic(err)
	}

	err = newGitHubClient().AssignIssue(repo, *ghPullRequest.Number, *issue.Assignee)

	if err != nil {
		panic(err)
	}

	return gh.toCommonPullRequest(ghPullRequest)
}

func (gh GitHub) AssignPullRequestReviewer(pr *PullRequest, reviewer string) {
	repo := versioncontrol.GetClient().Repo()

	err := newGitHubClient().RequestReviewer(repo, pr.Number, reviewer)

	if err != nil {
		panic(err)
	}
}

func (gh GitHub) MergePullRequest(pr *PullRequest, mergeMethod string) bool {
	repo := versioncontrol.GetClient().Repo()

	merged, err := newGitHubClient().MergePullRequest(repo, pr.Number, mergeMethod)

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
