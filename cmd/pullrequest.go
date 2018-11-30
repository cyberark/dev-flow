package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/scm"
	"github.com/cyberark/dev-flow/util"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var pullrequestCmd = &cobra.Command{
	Use:     "pullrequest",
	Aliases: []string{"pr"},
	Short:   "Creates a pull request for your branch.",
	Run: func(cmd *cobra.Command, args []string) {
		branchName := versioncontrol.GetClient().CurrentBranch()

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
			issue, err := issuetracking.GetClient().Issue(issueKey)

			if err != nil {
				log.Fatalln(err)
			}
			
			pr = scm.CreatePullRequest(issue)
		}

		if util.Confirm("Open pull request in browser?") {
			util.Openbrowser(pr.URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullrequestCmd)
}
