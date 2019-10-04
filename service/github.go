package service

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
)

type GitHubService interface {
	GetUser(string) (*github.User, error)
	GetIssues(common.Repo) ([]*github.Issue, error)
	GetIssue(repo common.Repo, issueNum int) (*github.Issue, error)
	AssignIssue(repo common.Repo, issueNum int, assigneeLogin string) (error)
	GetLabel(repo common.Repo, name string) (*github.Label, error)
	AddLabelToIssue(repo common.Repo, issueNum int, labelName string) (error)
	RemoveLabelForIssue(repo common.Repo, issueNum int, labelName string) (error)
	GetPullRequests(repo common.Repo, branchName string) ([]*github.PullRequest, error)
	GetPullRequest(repo common.Repo, pullRequestNum int) (*github.PullRequest, error)
}

type GitHub struct{}

func newClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}

func (gh GitHub) GetUser(login string) (*github.User, error) {
	user, _, err := newClient().Users.Get(context.Background(), login)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (gh GitHub) GetIssues(repo common.Repo) ([]*github.Issue, error) {
	issues, _, err := newClient().Issues.ListByRepo(
		context.Background(),
		repo.Owner,
		repo.Name,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (gh GitHub) GetIssue(repo common.Repo, issueNum int) (*github.Issue, error) {
	issue, _, err := newClient().Issues.Get(
		context.Background(),
		repo.Owner,
		repo.Name,
		issueNum,
	)

	if err != nil {
		return nil, err
	}

	return issue, nil
}

func (gh GitHub) AssignIssue(repo common.Repo, issueNum int, assigneeLogin string) (error) {
	_, _, err := newClient().Issues.AddAssignees(
		context.Background(),
		repo.Owner,
		repo.Name,
		issueNum,
		[]string{assigneeLogin},
	)

	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) GetLabel(repo common.Repo, name string) (*github.Label, error) {
	label, _, err := newClient().Issues.GetLabel(
		context.Background(),
		repo.Owner,
		repo.Name,
		name,
	)

	if err != nil {
		return nil, err
	}

	return label, nil
}

func (gh GitHub) AddLabelToIssue(repo common.Repo, issueNum int, labelName string) (error) {
	_, _, err := newClient().Issues.AddLabelsToIssue(
		context.Background(),
		repo.Owner,
		repo.Name,
		issueNum,
		[]string{labelName},
	)

	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) RemoveLabelForIssue(repo common.Repo, issueNum int, labelName string) (error) {
	_, err := newClient().Issues.RemoveLabelForIssue(
		context.Background(),
		repo.Owner,
		repo.Name,
		issueNum,
		labelName,
	)

	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) GetPullRequests(repo common.Repo, branchName string) ([]*github.PullRequest, error) {
	opts := &github.PullRequestListOptions{
		State: "open",
		Base:  "master",
		Head:  fmt.Sprintf("%v:%v", repo.Owner, branchName),
	}

	pullRequests, _, err := newClient().PullRequests.List(
		context.Background(),
		repo.Owner,
		repo.Name,
		opts,
	)

	if err != nil {
		return nil, err
	}

	return pullRequests, nil
}

func (gh GitHub) GetPullRequest(repo common.Repo, pullRequestNum int) (*github.PullRequest, error) {
	pullRequest, _, err := newClient().PullRequests.Get(
		context.Background(),
		repo.Owner,
		repo.Name,
		pullRequestNum,
	)

	if err != nil {
		return nil, err
	}

	return pullRequest, nil
}

func (gh GitHub) CreatePullRequest(repo common.Repo, newPullRequest *github.NewPullRequest) (*github.PullRequest, error) {
	pullRequest, _, err := newClient().PullRequests.Create(
		context.Background(),
		repo.Owner,
		repo.Name,
		newPullRequest,
	)

	if err != nil {
		return nil, err
	}

	return pullRequest, nil
}

func (gh GitHub) RequestReviewer(repo common.Repo, pullRequestNum int, reviewerLogin string) (error) {
	reviewersRequest := github.ReviewersRequest{
		Reviewers: []string{reviewerLogin},
	}
	
	_, _, err := newClient().PullRequests.RequestReviewers(
		context.Background(),
		repo.Owner,
		repo.Name,
		pullRequestNum,
		reviewersRequest,
	)

	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) MergePullRequest(repo common.Repo, pullRequestNum int, mergeMethod string) (bool, error) {
	pullRequestOptions := &github.PullRequestOptions{
		MergeMethod: mergeMethod,
	}

	mergeResult, _, err := newClient().PullRequests.Merge(
		context.Background(),
		repo.Owner,
		repo.Name,
		pullRequestNum,
		"",
		pullRequestOptions,
	)

	if err != nil {
		return false, err
	}

	return *mergeResult.Merged, nil
}
