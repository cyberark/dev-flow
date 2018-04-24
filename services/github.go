package services

import (
	"context"
	"golang.org/x/oauth2"

	"github.com/spf13/viper"
	
	"github.com/google/go-github/github"

	"github.com/jtuttle/dev-flow/common"
)

type GitHub struct { }

func (gh GitHub) GetClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{ AccessToken: viper.Get("github_access_token").(string) },
	)
	tc := oauth2.NewClient(context.Background(), ts)
	
	return github.NewClient(tc)
}

func (gh GitHub) ToCommonIssue(ghIssue *github.Issue) common.Issue {
	var assignee *string

	if ghIssue.Assignee != nil {
		assignee = ghIssue.Assignee.Login
	}
	
	ghLabels := ghIssue.Labels
	labels := make([]string, len(ghLabels))

	for i, ghLabel := range ghLabels {
		labels[i] = *ghLabel.Name
	}
	
	return common.Issue {
		URL: ghIssue.HTMLURL,
		Number: ghIssue.Number,
		Title: ghIssue.Title,
		Assignee: assignee,
		Labels: labels,
	}
}
