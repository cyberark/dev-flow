package scm

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	
	"github.com/conjurinc/dev-flow/common"
	"github.com/conjurinc/dev-flow/services"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

type GitHub struct{}

func (gh GitHub) client() *github.Client {
	return services.GitHub{}.GetClient()
}

func (gh GitHub) GetPullRequest(branchName string) *PullRequest {
	repo := versioncontrol.GetClient().Repo()

	client := gh.client()

	opts := &github.PullRequestListOptions {
		State: "open",
		Base: "master",
		Head: fmt.Sprintf("%v:%v", repo.Owner, branchName),
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

func (gh GitHub) CreatePullRequest(issue common.Issue) *PullRequest {
	repo := versioncontrol.GetClient().Repo()
	
	base := "master"
	head := issue.BranchName()
	
	ghnpr := &github.NewPullRequest {
		Title: issue.Title,
		Base: &base,
		Head: &head,
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

	return gh.toCommonPullRequest(ghpr)
}

func (gh GitHub) MergePullRequest(pr *PullRequest) bool {
	repo := versioncontrol.GetClient().Repo()
	
	client := gh.client()
	
	ghmr, _, err := client.PullRequests.Merge(
		context.Background(),
		repo.Owner,
		repo.Name,
		pr.Number,
		"",
		nil,
	)

	if err != nil {
		panic(err)
	}
	
	return *ghmr.Merged
}

func (gh GitHub) toCommonPullRequest(ghpr *github.PullRequest) *PullRequest {
	return &PullRequest {
		Number: *ghpr.Number,
		Creator: *ghpr.User.Login,
		Base: *ghpr.Base.Ref,
		Mergeable: *ghpr.Mergeable,
		URL: *ghpr.HTMLURL,
	}
}
