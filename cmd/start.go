package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

		user := it.GetCurrentUser()
		it.AssignIssue(issue, user)
		fmt.Printf("Assigned issue %v to user %v.\n", *issue.Number, user)

		progressLabelName := viper.Get("labels.in_progress")

		if progressLabelName != nil {
			err := it.AddIssueLabel(issue, progressLabelName.(string))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		
			fmt.Printf("Added label '%v' to issue %v.\n", progressLabelName, *issue.Number)
		}

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
