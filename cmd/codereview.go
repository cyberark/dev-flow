package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conjurinc/dev-flow/chat"
	"github.com/conjurinc/dev-flow/issuetracking"
	"github.com/conjurinc/dev-flow/scm"
	"github.com/conjurinc/dev-flow/util"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

var codereviewCmd = &cobra.Command{
	Use:   "codereview [reviewer]",
	Aliases: []string { "cr" },
	Short: "Creates a pull request and assigns a reviewer.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reviewer := args[0]
		
		branchName := versioncontrol.GetClient().CurrentBranch()

		it := issuetracking.GetClient()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue := it.Issue(issueKey)

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)
		
		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			pr = scm.CreatePullRequest(issue)
		}

		it.AssignIssue(issue, reviewer)

		chat := chat.GetClient()
		chat.DirectMessage(
			reviewer,
			fmt.Sprintf("%v has assigned you as reviewer on %v", it.GetCurrentUser(), pr.URL),
		)
		
		if util.Confirm("Open pull request in browser?") {
			util.Openbrowser(pr.URL)	
		}
	},
}

func init() {
	rootCmd.AddCommand(codereviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codereviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codereviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
