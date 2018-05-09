package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	
	"github.com/conjurinc/dev-flow/issuetracking"
	"github.com/conjurinc/dev-flow/scm"
	"github.com/conjurinc/dev-flow/util"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

var pullrequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Aliases: []string { "pr" },
	Short: "Creates a pull request for your branch.",
	Run: func(cmd *cobra.Command, args []string) {
		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			it := issuetracking.GetClient()
			issue := it.Issue(issueKey)
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
