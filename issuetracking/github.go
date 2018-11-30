package issuetracking

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/services"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHub struct{}

func (gh GitHub) client() *github.Client {
	return services.GitHub{}.GetClient()
}

func (gh GitHub) getUser(username string) *github.User {
	client := gh.client()
	ghUser, _, err := client.Users.Get(context.Background(), username)

	if err != nil {
		panic(err)
	}

	return ghUser
}

func (gh GitHub) GetCurrentUser() string {
	return *gh.getUser("").Login
}

func (gh GitHub) GetUserRealName(username string) string {
	return *gh.getUser(username).Name
}

func (gh GitHub) Issues() []common.Issue {
	repo := versioncontrol.Git{}.Repo()

	ghIssues, _, err := gh.client().Issues.ListByRepo(
		context.Background(),
		repo.Owner,
		repo.Name,
		nil,
	)

	if err != nil {
		panic(err)
	}

	var issues []common.Issue

	for _, ghIssue := range ghIssues {
		if ghIssue.PullRequestLinks == nil {
			issue := services.GitHub{}.ToCommonIssue(ghIssue)
			issues = append(issues, issue)
		}
	}

	return issues
}

func (gh GitHub) Issue(issueKey string) (common.Issue, error) {
	repo := versioncontrol.Git{}.Repo()

	client := gh.client()

	issueNum, err := strconv.Atoi(issueKey)

	if err != nil {
		panic(err)
	}

	ghIssue, _, err := client.Issues.Get(
		context.Background(),
		repo.Owner,
		repo.Name,
		issueNum,
	)

	if err != nil {
		return common.Issue{},
		errors.New(fmt.Sprintf("Could not find issue number %d", issueNum))
	}

	return services.GitHub{}.ToCommonIssue(ghIssue), nil
}

func (gh GitHub) AssignIssue(issue common.Issue, login string) {
	repo := versioncontrol.Git{}.Repo()

	client := gh.client()

	_, _, err := client.Issues.AddAssignees(
		context.Background(),
		repo.Owner,
		repo.Name,
		*issue.Number,
		[]string{login},
	)

	if err != nil {
		panic(err)
	}
}

func (gh GitHub) AddIssueLabel(issue common.Issue, labelName string) error {
	repo := versioncontrol.Git{}.Repo()

	client := gh.client()

	_, err := gh.getLabel(labelName)

	if err != nil {
		return errors.New(fmt.Sprintf("Label '%v' does not exist.", labelName))
	}

	labels := []string{labelName}

	_, _, err = client.Issues.AddLabelsToIssue(
		context.Background(),
		repo.Owner,
		repo.Name,
		*issue.Number,
		labels,
	)

	if err != nil {
		return err
	}

	fmt.Println("Added label '%s' to issue %d", labelName, issue.Number)

	return nil
}

func (gh GitHub) RemoveIssueLabel(issue common.Issue, labelName string) {
	repo := versioncontrol.Git{}.Repo()

	client := gh.client()

	_, err := client.Issues.RemoveLabelForIssue(
		context.Background(),
		repo.Owner,
		repo.Name,
		*issue.Number,
		labelName,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Removed label '%s' from issue %d", labelName, issue.Number)
}

func (gh GitHub) getLabel(name string) (*github.Label, error) {
	repo := versioncontrol.GetClient().Repo()

	client := gh.client()

	ghl, _, err := client.Issues.GetLabel(
		context.Background(),
		repo.Owner,
		repo.Name,
		name,
	)

	if err != nil {
		return nil, err
	}

	return ghl, nil
}
