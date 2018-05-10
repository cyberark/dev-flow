package issuetracking

import (
	"context"
	"strconv"
	
	"github.com/google/go-github/github"

	"github.com/conjurinc/dev-flow/common"
	"github.com/conjurinc/dev-flow/services"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

type GitHub struct{}

func (gh GitHub) client() *github.Client {
	return services.GitHub{}.GetClient()
}

func (gh GitHub) GetCurrentUser() string {
	client := gh.client()
	ghuser, _, err := client.Users.Get(context.Background(), "")

	if err != nil {
		panic(err)
	}

	return *ghuser.Login
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

	issues := make([]common.Issue, len(ghIssues))

	for i, ghIssue := range ghIssues {
		issues[i] = services.GitHub{}.ToCommonIssue(ghIssue)
	}

	return issues
}

func (gh GitHub) Issue(issueKey string) common.Issue {
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
		panic(err)
	}

	return services.GitHub{}.ToCommonIssue(ghIssue)
}

func (gh GitHub) AssignIssue(issue common.Issue, login string) {
	repo := versioncontrol.Git{}.Repo()
	
	client := gh.client()

	_, _, err := client.Issues.AddAssignees(
		context.Background(),
		repo.Owner,
		repo.Name,
		*issue.Number,
		[]string { login },
	)

	if err != nil {
		panic(err)
	}
}
