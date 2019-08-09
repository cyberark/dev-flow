package service

import (
	"encoding/json"
	"fmt"
	
	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/testutils"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHubMock struct{
	Error error
}

func (ghm GitHubMock) GetUser(login string) (*github.User, error) {
	fmt.Println("****")
	fmt.Println(ghm.Error)
	fmt.Println("****")
	
	if ghm.Error != nil {
		return nil, ghm.Error
	}
	
	// Parent folder will be wherever tests are run from.
	userBytes := testutils.LoadJsonFile("testdata/github_user.json")

	var user github.User
	err := json.Unmarshal(userBytes, &user)

	if err != nil {
		fmt.Println("There was an error:", err)
	}
	
	return &user, nil
}

func (ghm GitHubMock) GetIssues(repo versioncontrol.Repo) ([]*github.Issue, error) {


	return nil, nil
}

func (ghm GitHubMock) GetIssue(repo versioncontrol.Repo, issueNum int) (*github.Issue, error) {


	return nil, nil
}

func (ghm GitHubMock) AssignIssue(repo versioncontrol.Repo, issueNum int, assigneeLogin string) (error) {

	
	return nil
}

func (ghm GitHubMock) GetLabel(repo versioncontrol.Repo, name string) (*github.Label, error) {

	
	return nil, nil
}

func (ghm GitHubMock) AddLabelToIssue(repo versioncontrol.Repo, issueNum int, labelName string) (error) {

	
	return nil
}

func (ghm GitHubMock) RemoveLabelForIssue(repo versioncontrol.Repo, issueNum int, labelName string) (error) {

	
	return nil
}

func (ghm GitHubMock) GetPullRequests(repo versioncontrol.Repo, branchName string) ([]*github.PullRequest, error) {

	
	return nil, nil
}

func (ghm GitHubMock) GetPullRequest(repo versioncontrol.Repo, pullRequestNum int) (*github.PullRequest, error) {

	
	return nil, nil
}

func (ghm GitHubMock) CreatePullRequest(repo versioncontrol.Repo, newPullRequest *github.NewPullRequest) (*github.PullRequest, error) {

	
	return nil, nil
}

func (ghm GitHubMock) RequestReviewer(repo versioncontrol.Repo, pullRequestNum int, reviewerLogin string) (error) {

	
	return nil
}

func (ghm GitHubMock) MergePullRequest(repo versioncontrol.Repo, pullRequestNum int, mergeMethod string) (bool, error) {

	
	return true, nil
}
