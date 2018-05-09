package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/conjurinc/dev-flow/issuetracking"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Creates a remote branch and initial commit for the specified issue",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		issueKey := args[0]
		
		it := issuetracking.GetClient()
		issue := it.Issue(issueKey)
		it.AssignIssue(issue, it.GetCurrentUser())

		vc := versioncontrol.GetClient()
		
		vc.CheckoutAndPull("master")

		branchName := issue.BranchName()
		
		if vc.IsRemoteBranch(branchName) {
			vc.CheckoutAndPull(branchName)
		} else {
			vc.InitBranch(*issue.Number, branchName)
		}

		fmt.Println("Issue started! You are now working in branch:", branchName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
