package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cyberark/dev-flow/chat"
	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/scm"
	"github.com/cyberark/dev-flow/util"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var MergeMethod string = "rebase"

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Squash merges the story branch and completes the issue.",
	Run: func(cmd *cobra.Command, args []string) {
		util.ValidateStringParam(
			"merge-method",
			MergeMethod,
			[]string{ "rebase", "squash", "merge" },
		)
		
		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr == nil {
			err := fmt.Sprintf("No pull request found for branch %s", branchName)
			log.Fatalln(err)
		}

		if !pr.Mergeable {
			err := "Pull request not mergeable. Check for conflicts."
			log.Fatalln(err)
		}

		if !util.Confirm(fmt.Sprintf("Are you sure you want to merge %v into %v?", branchName, pr.Base)) {
			fmt.Println("Pull request not merged.")
			os.Exit(0)
		}

		success := scm.MergePullRequest(pr, MergeMethod)

		it := issuetracking.GetClient()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue, err := it.Issue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		if success {
			fmt.Printf("Merged %v into %v\n", branchName, pr.Base)
		} else {
			err := "Merge failed"
			log.Fatalln(err)
		}

		it.AssignIssue(issue, pr.Creator)

		chat := chat.GetClient()

		if chat != nil {
			userRealName, err := it.GetUserRealName(pr.Creator)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			userLogin, err := it.GetCurrentUser()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			
			chat.DirectMessage(
				userRealName,
				fmt.Sprintf("%v has merged your pull request %v", userLogin, pr.URL),
			)
		}

		err = it.RemoveIssueLabel(issue, viper.GetString("labels.codereview"))

		if err != nil {
			log.Println(err)
		}
		
		err = it.AddIssueLabel(issue, viper.GetString("labels.complete"))

		if err != nil {
			log.Println(err)
		}

		vc.CheckoutAndPull(pr.Base)

		if util.Confirm(fmt.Sprintf("Delete remote branch %v", branchName)) {
			vc.DeleteRemoteBranch(branchName)
			fmt.Println("Remote branch deleted.")
		}

		if util.Confirm(fmt.Sprintf("Delete local branch %v", branchName)) {
			vc.DeleteLocalBranch(branchName)
			fmt.Println("Local branch deleted.")
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	
	completeCmd.Flags().StringVarP(
		&MergeMethod,
		"merge-method",
		"m",
		"rebase",
		"Merge method to use (rebase, squash, or merge). Defaults to rebase.",
	)
}
