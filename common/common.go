package common

import (
	"fmt"
	"regexp"
	"strings"
)

type Issue struct {
	URL string
	Number int
	Title string
	Assignee string
	Labels []string
}

func (issue Issue) String() string {
	return fmt.Sprintf("%v - %v", issue.Number, issue.Title)
}

func (issue Issue) BranchName() string {
	title := issue.Title
	title = strings.ToLower(title)
	title = strings.TrimSpace(title)

	re := regexp.MustCompile(`[^\w\s]`)
	title = re.ReplaceAllString(title, "$1$1")

	title = strings.Replace(title, " ", "-", -1)
	
	return fmt.Sprintf("%v--%v", issue.Number, title)
}

func (issue Issue) HasLabel(label string) bool {
	for _, issueLabel := range issue.Labels {
		if issueLabel == label {
			return true
		}
	}
	return false
}
