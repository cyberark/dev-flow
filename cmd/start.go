package cmd

import (
	"fmt"
	"log"
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
		issue, err := it.GetIssue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		login, err := it.GetCurrentUserLogin()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		err = it.AssignIssue(issue.Number, login)

		if err != nil {
			log.Fatalln(err)
		}
		
		fmt.Printf("Assigned issue %v to user %v.\n", issue.Number, login)

		err = it.AddIssueLabel(*issue, viper.GetString("labels.start"))

		if err != nil {
			log.Println(err)
		}

		vc := versioncontrol.GetClient()

		vc.CheckoutAndPull("master")

		branchName := issue.BranchName()

		if vc.IsRemoteBranch(branchName) {
			vc.CheckoutAndPull(branchName)
		} else {

			vc.InitBranch(issue.Number, branchName)
		}

		fmt.Println("Issue started! You are now working in branch:", branchName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
