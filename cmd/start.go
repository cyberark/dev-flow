package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Creates a remote branch for the specified issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		issueKey := args[0]

		it := issuetracking.GetClient()
		issue, err := it.Issue(issueKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		user := it.GetCurrentUser()
		it.AssignIssue(issue, user)
		fmt.Printf("Assigned issue %v to user %v.\n", *issue.Number, user)

		progressLabelName := viper.GetString("labels.start")

		if progressLabelName != "" {
			err := it.AddIssueLabel(issue, progressLabelName)

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
