package issuetracking

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHub struct{}

func newGitHubClient() service.GitHub {
	return service.GitHub{}.GetClient()
}

func toCommonIssue(ghIssue *github.Issue) common.Issue {
	var assignee *string

	if ghIssue.Assignee != nil {
		assignee = ghIssue.Assignee.Login
	}

	ghLabels := ghIssue.Labels
	labels := make([]string, len(ghLabels))

	for i, ghLabel := range ghLabels {
		labels[i] = *ghLabel.Name
	}

	return common.Issue{
		URL:      ghIssue.HTMLURL,
		Number:   ghIssue.Number,
		Title:    ghIssue.Title,
		Assignee: assignee,
		Labels:   labels,
	}
}

func (gh GitHub) getUser(username string) *github.User {
	ghUser, err := newGitHubClient().GetUser(username)

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

	ghIssues, err := newGitHubClient().GetIssues(repo)

	if err != nil {
		panic(err)
	}

	var issues []common.Issue

	for _, ghIssue := range ghIssues {
		if ghIssue.PullRequestLinks == nil {
			issue := toCommonIssue(ghIssue)
			issues = append(issues, issue)
		}
	}

	return issues
}

func (gh GitHub) Issue(issueKey string) (common.Issue, error) {
	repo := versioncontrol.Git{}.Repo()

	issueNum, err := strconv.Atoi(issueKey)

	if err != nil {
		panic(err)
	}

	ghIssue, err := newGitHubClient().GetIssue(repo, issueNum)

	if err != nil {
		return common.Issue{},
		errors.New(fmt.Sprintf("Could not find issue number %d", issueNum))
	}

	return toCommonIssue(ghIssue), nil
}

func (gh GitHub) AssignIssue(issue common.Issue, login string) {
	repo := versioncontrol.Git{}.Repo()

	err := newGitHubClient().AssignIssue(repo, *issue.Number, login)
	
	if err != nil {
		panic(err)
	}
}

func (gh GitHub) AddIssueLabel(issue common.Issue, labelName string) error {
	if labelName == "" {
		return errors.New("Unable to add blank label.")
	}

	if issue.HasLabel(labelName) {
		return fmt.Errorf("Issue %d already has label '%s'.", *issue.Number, labelName)
	}

	_, err := gh.getLabel(labelName)

	if err != nil {
		return err
	}

	repo := versioncontrol.Git{}.Repo()

	err = newGitHubClient().AddLabelToIssue(repo, *issue.Number, labelName)
	
	if err != nil {
		return fmt.Errorf("Failed to add label '%s' to issue %d: %s", labelName, *issue.Number, err)
	}

	return nil
}

func (gh GitHub) RemoveIssueLabel(issue common.Issue, labelName string) error {
	if labelName == "" {
		return errors.New("Unable to remove blank label.")
	}

	if !issue.HasLabel(labelName) {
		return fmt.Errorf("Issue %d does not have label '%s'.", *issue.Number, labelName)
	}
	
	repo := versioncontrol.Git{}.Repo()
	
	err := newGitHubClient().RemoveLabelForIssue(repo, *issue.Number, labelName)

	if err != nil {
		return fmt.Errorf("Failed to remove label '%s' from issue %d: %s", labelName, *issue.Number, err)
	}
	
	return nil
}

func (gh GitHub) getLabel(name string) (*github.Label, error) {
	repo := versioncontrol.GetClient().Repo()

	ghLabel, err := newGitHubClient().GetLabel(repo, name)

	if err != nil {
		return nil, err
	}

	return ghLabel, nil
}
