package cmd

import (
	"fmt"
	"log"

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

		vc := versioncontrol.GetClient()
		repo, err := vc.Repo()

		if err != nil {
			log.Fatalln(err)
		}

		it := issuetracking.GetClient(repo)
		issue, err := it.GetIssue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		login, err := it.GetCurrentUserLogin()

		if err != nil {
			log.Fatalln(err)
		}
		
		err = it.AssignIssue(issue.Number, login)

		if err != nil {
			log.Fatalln(err)
		}
		
		fmt.Printf("Assigned issue %v to user %v.\n", issue.Number, login)

		err = it.AddIssueLabel(issue.Number, viper.GetString("labels.start"))

		if err != nil {
			log.Println(err)
		}

		vc.CheckoutAndPull("master")

		branchName := issue.BranchName()

		isRemote, err := vc.IsRemoteBranch(branchName)

		if err != nil {
			log.Fatalln(err)
		}
		
		if isRemote {
			output, err := vc.CheckoutAndPull(branchName)

			if err != nil {
				log.Fatalln(err)
			}
			
			log.Println(output)			
		} else {
			output, err := vc.InitBranch(issue.Number, branchName)
			
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(output)
		}

		fmt.Println("Issue started! You are now working in branch:", branchName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
