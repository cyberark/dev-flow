package cmd

import (
	"fmt"
	"log"

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
		it := issuetracking.GetClient(vc.Repo())
		issue, err := it.GetIssue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}
		
		err = it.AssignIssue(issue.Number, pr.Creator)

		if err != nil {
			log.Fatalln(err)
		}

		chat := chat.GetClient()

		if chat != nil {
			userRealName, err := it.GetUserRealName(pr.Creator)

			if err != nil {
				log.Fatalln(err)
			}

			login, err := it.GetCurrentUserLogin()

			if err != nil {
				log.Fatalln(err)
			}
			
			chat.DirectMessage(
				userRealName,
				fmt.Sprintf("%v has requested changes on %v", login, pr.URL),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviseCmd)
}
