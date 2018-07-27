package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/conjurinc/dev-flow/issuetracking"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List a user's GitHub issues.",
	Long: "Lists issues assigned to a user across all projects.",
	Run: func(cmd *cobra.Command, args []string) {
		it := issuetracking.GetClient()

		issues := it.Issues()

		for _, issue := range issues {
			assignee := "unassigned"
			
			if issue.Assignee != nil {
				assignee = *issue.Assignee
			}
			
			fmt.Printf(
				"%v - %v (%v) [%v] \n",
				*issue.Number,
				*issue.Title,
				assignee,
				strings.Join(issue.Labels, ", "),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}
