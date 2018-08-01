package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/conjurinc/dev-flow/issuetracking"
	"github.com/conjurinc/dev-flow/scm"
	"github.com/conjurinc/dev-flow/util"
	"github.com/conjurinc/dev-flow/versioncontrol"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Merges the story branch and completes the issue.",
	Run: func(cmd *cobra.Command, args []string) {		
		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr == nil {
			fmt.Println("No pull request found for branch", branchName)
			os.Exit(1)
		}
		
		if !pr.Mergeable {
			fmt.Println("Pull request not mergeable. Check for conflicts.")
			os.Exit(1)
		}

		if !util.Confirm(fmt.Sprintf("Are you sure you want to merge %v into %v", branchName, pr.Base)) {
			fmt.Println("Pull request not merged.")
			os.Exit(0)
		}

		success := scm.MergePullRequest(pr)

		it := issuetracking.GetClient()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue := it.Issue(issueKey)

		if success {
			fmt.Printf("Merged %v into %v\n", branchName, pr.Base)
		} else {
			fmt.Println("Merge failed.")
			os.Exit(1)
		}

		it.AssignIssue(issue, pr.Creator)

		reviewLabelName := viper.GetString("labels.in_review")
		
		if reviewLabelName != "" && issue.HasLabel(reviewLabelName) {
			it.RemoveIssueLabel(issue, reviewLabelName)
			fmt.Printf("Removed label '%v' from issue %v.\n", reviewLabelName, *issue.Number)
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
}
