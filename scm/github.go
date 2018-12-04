package scm

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/services"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHub struct{}

func (gh GitHub) client() *github.Client {
	return services.GitHub{}.GetClient()
}

func (gh GitHub) GetPullRequest(branchName string) *PullRequest {
	repo := versioncontrol.GetClient().Repo()

	client := gh.client()

	opts := &github.PullRequestListOptions{
		State: "open",
		Base:  "master",
		Head:  fmt.Sprintf("%v:%v", repo.Owner, branchName),
	}

	ghprs, _, err := client.PullRequests.List(
		context.Background(),
		repo.Owner,
		repo.Name,
		opts,
	)

	if err != nil {
		panic(err)
	}

	var prnum *int

	for _, ghpr := range ghprs {
		if *ghpr.Head.Ref == branchName {
			prnum = ghpr.Number
		}
	}

	var pr *PullRequest

	if prnum != nil {
		// We call List and then Get because the Mergeable field is only
		// returned when retrieving single issues.
		ghpr, _, err := client.PullRequests.Get(
			context.Background(),
			repo.Owner,
			repo.Name,
			*prnum,
		)

		if err != nil {
			panic(err)
		}

		pr = gh.toCommonPullRequest(ghpr)
	}

	return pr
}

func (gh GitHub) CreatePullRequest(issue common.Issue, linkType string) *PullRequest {
	repo := versioncontrol.GetClient().Repo()

	base := "master"
	head := issue.BranchName()
	body := fmt.Sprintf("%s #%v", linkType, *issue.Number)

	ghnpr := &github.NewPullRequest{
		Base:  &base,
		Head:  &head,
		Title: issue.Title,
		Body:  &body,
	}

	client := gh.client()

	ghpr, _, err := client.PullRequests.Create(
		context.Background(),
		repo.Owner,
		repo.Name,
		ghnpr,
	)

	if err != nil {
		panic(err)
	}

	_, _, err = client.Issues.AddAssignees(
		context.Background(),
		repo.Owner,
		repo.Name,
		*ghpr.Number,
		[]string{*issue.Assignee},
	)

	if err != nil {
		panic(err)
	}

	return gh.toCommonPullRequest(ghpr)
}

func (gh GitHub) AssignPullRequestReviewer(pr *PullRequest, reviewer string) {
	repo := versioncontrol.GetClient().Repo()

	client := gh.client()

	reviewers := github.ReviewersRequest{
		Reviewers: []string{reviewer},
	}

	_, _, err := client.PullRequests.RequestReviewers(
		context.Background(),
		repo.Owner,
		repo.Name,
		pr.Number,
		reviewers,
	)

	if err != nil {
		panic(err)
	}
}

func (gh GitHub) MergePullRequest(pr *PullRequest, mergeMethod string) bool {
	repo := versioncontrol.GetClient().Repo()

	client := gh.client()

	prOpt := &github.PullRequestOptions{
		MergeMethod: mergeMethod,
	}

	ghmr, _, err := client.PullRequests.Merge(
		context.Background(),
		repo.Owner,
		repo.Name,
		pr.Number,
		"",
		prOpt,
	)

	if err != nil {
		panic(err)
	}

	return *ghmr.Merged
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
