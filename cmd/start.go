package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/jtuttle/dev-flow/issuetracking"
	"github.com/jtuttle/dev-flow/versioncontrol"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
