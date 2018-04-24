package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"

	"github.com/jtuttle/dev-flow/scm"
	"github.com/jtuttle/dev-flow/util"
	"github.com/jtuttle/dev-flow/versioncontrol"
)

// completeCmd represents the complete command
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

		if success {
			fmt.Println("Merged %v into %v", branchName, pr.Base)
		} else {
			fmt.Println("Merge failed.")
			os.Exit(1)
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

		// assign and close issue
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
