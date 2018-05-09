package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/conjurinc/dev-flow/scm"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

var reviseCmd = &cobra.Command{
	Use:   "revise",
	Short: "Rejects a PR and assigns it back to the implementor.",
	Run: func(cmd *cobra.Command, args []string) {
		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()
		
		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		// TODO: This won't work when the issue tracker != the scm
		// for example Jira vs GitHub
		it := issuetracking.GetClient()
		issue := it.Issue(issueKey)
		it.AssignIssue(issue, pr.Creator)

		// Notify user in Slack?
	},
}

func init() {
	rootCmd.AddCommand(reviseCmd)
}