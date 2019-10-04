package cmd

import (
	"fmt"
	"log"

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
			log.Fatalln("Pull request not merged.")
		}

		success := scm.MergePullRequest(pr, MergeMethod)

		it := issuetracking.GetClient(vc.Repo())
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue, err := it.GetIssue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		if success {
			fmt.Printf("Merged %v into %v\n", branchName, pr.Base)
		} else {
			err := "Merge failed"
			log.Fatalln(err)
		}

		err = it.AssignIssue(issue.Number, pr.Creator)

		if err != nil {
			log.Fatalln(err)
		}
		
		chat := chat.GetClient()

		if chat != nil {
			userRealName, err := it.GetUserRealName(pr.Creator)

			if err != nil {
				log.Fatalln(err)
			}

			login, err := it.GetCurrentUserLogin()

			if err != nil {
				log.Fatalln(err)
			}
			
			chat.DirectMessage(
				userRealName,
				fmt.Sprintf("%v has merged your pull request %v", login, pr.URL),
			)
		}

		err = it.RemoveIssueLabel(issue.Number, viper.GetString("labels.codereview"))

		if err != nil {
			log.Println(err)
		}
		
		err = it.AddIssueLabel(issue.Number, viper.GetString("labels.complete"))

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
