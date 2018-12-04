package cmd

import (
	"fmt"
	"log"
	
	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var IssueKey string

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Add a label to an issue.",
	Long:  "Apply a label to an issue.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		label := args[0]

		if IssueKey == "" {
			fmt.Println("No issue key provided, retrieving from branch.")
			branchName := versioncontrol.GetClient().CurrentBranch()
			IssueKey = issuetracking.GetIssueKeyFromBranchName(branchName)
		}

		if IssueKey == "" {
			log.Fatalln("No issue key provided")
		}
		
		it := issuetracking.GetClient()
		issue, err := it.Issue(IssueKey)

		if err != nil {
			log.Fatalln(err)
		}
		
		it.AddIssueLabel(issue, label)
	},
}

func init() {
	rootCmd.AddCommand(labelCmd)
	
	labelCmd.Flags().StringVarP(
		&IssueKey,
		"issue-key",
		"i",
		"",
		"The key of the issue to which the label should be added.",
	)
}
