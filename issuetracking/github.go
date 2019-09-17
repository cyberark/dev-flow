package issuetracking

import (
	"strconv"

	"github.com/google/go-github/github"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
	"github.com/cyberark/dev-flow/versioncontrol"
)

type GitHub struct {
	GitHubService service.GitHubService
}

func toCommonIssue(ghIssue *github.Issue) common.Issue {
	var assignee string = ""

	if ghIssue.Assignee != nil {
		assignee = *ghIssue.Assignee.Login
	}

	ghLabels := ghIssue.Labels
	labels := make([]string, len(ghLabels))

	for i, ghLabel := range ghLabels {
		labels[i] = *ghLabel.Name
	}

	return common.Issue{
		URL:      *ghIssue.HTMLURL,
		Number:   *ghIssue.Number,
		Title:    *ghIssue.Title,
		Assignee: assignee,
		Labels:   labels,
	}
}

func (gh GitHub) getUser(login string) (*github.User, error) {
	ghUser, err := gh.GitHubService.GetUser(login)
	return ghUser, err
}

func (gh GitHub) GetCurrentUserLogin() (string, error) {
	ghUser, err := gh.getUser("")

	if err != nil {
		return "", err
	}
	
	return *ghUser.Login, nil
}

func (gh GitHub) GetUserRealName(login string) (string, error) {
	ghUser, err := gh.getUser(login)

	if err != nil {
		return "", err
	}
	
	return *ghUser.Name, nil
}

func (gh GitHub) GetIssues() ([]common.Issue, error) {
	repo := versioncontrol.Git{}.Repo()

	ghIssues, err := gh.GitHubService.GetIssues(repo)
	
	if err != nil {
		return nil, err
	}

	var issues []common.Issue

	for _, ghIssue := range ghIssues {
		issue := toCommonIssue(ghIssue)
		issues = append(issues, issue)
	}

	return issues, nil
}

func (gh GitHub) GetIssue(issueKey string) (*common.Issue, error) {
	repo := versioncontrol.Git{}.Repo()

	issueNum, err := strconv.Atoi(issueKey)

	if err != nil {
		return nil, err
	}

	ghIssue, err := gh.GitHubService.GetIssue(repo, issueNum)

	if err != nil {
		return nil, err
	}

	issue := toCommonIssue(ghIssue)

	return &issue, nil
}

func (gh GitHub) AssignIssue(issueNum int, login string) error {
	repo := versioncontrol.Git{}.Repo()

	err := gh.GitHubService.AssignIssue(repo, issueNum, login)
	
	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) AddIssueLabel(issueNum int, labelName string) error {
	repo := versioncontrol.Git{}.Repo()

	err := gh.GitHubService.AddLabelToIssue(repo, issueNum, labelName)
	
	if err != nil {
		return err
	}

	return nil
}

func (gh GitHub) RemoveIssueLabel(issueNum int, labelName string) error {
	repo := versioncontrol.Git{}.Repo()
	
	err := gh.GitHubService.RemoveLabelForIssue(repo, issueNum, labelName)

	if err != nil {
		return err
	}
	
	return nil
}
