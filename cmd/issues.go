package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/issuetracking"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Lists open, unassigned issues on the current repository.",
	Long:  "Lists open, unassigned issues on the current repository.",
	Run: func(cmd *cobra.Command, args []string) {
		it := issuetracking.GetClient()

		issues, err := it.Issues()

		if err != nil {
			log.Fatalln(err)
		}

		for _, issue := range issues {
			assignee := "unassigned"

			if issue.Assignee != "" {
				assignee = issue.Assignee
			}

			fmt.Printf(
				"%v - %v (%v) [%v] \n",
				issue.Number,
				issue.Title,
				assignee,
				strings.Join(issue.Labels, ", "),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}
