package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/chat"
	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/scm"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var reviseCmd = &cobra.Command{
	Use:   "revise",
	Short: "Rejects a PR and assigns it back to the implementor.",
	Run: func(cmd *cobra.Command, args []string) {
		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		// TODO: This won't work when the issue tracker != the scm
		// for example Jira vs GitHub
		it := issuetracking.GetClient()
		issue, err := it.Issue(issueKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		it.AssignIssue(issue, pr.Creator)

		chat := chat.GetClient()

		if chat != nil {
			chat.DirectMessage(
				it.GetUserRealName(pr.Creator),
				fmt.Sprintf("%v has requested changes on %v", it.GetCurrentUser(), pr.URL),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviseCmd)
}
