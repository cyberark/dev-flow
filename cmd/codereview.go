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

var codereviewCmd = &cobra.Command{
	Use:     "codereview [reviewer]",
	Aliases: []string{"cr"},
	Short:   "Creates a pull request and assigns a reviewer.",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		util.ValidateStringParam(
			"link-type",
			LinkType,
			[]string{ "close", "connect" },
		)
		
		reviewer := args[0]

		branchName := versioncontrol.GetClient().CurrentBranch()

		it := issuetracking.GetClient()
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue, err := it.Issue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		progressLabelName := viper.GetString("labels.start")

		if progressLabelName != "" && issue.HasLabel(progressLabelName) {
			it.RemoveIssueLabel(issue, progressLabelName)
			fmt.Printf("Removed label '%v' from issue %v.\n", progressLabelName, *issue.Number)
		}

		reviewLabelName := viper.GetString("labels.codereview")

		if reviewLabelName != "" {
			err := it.AddIssueLabel(issue, reviewLabelName)

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("Added label '%v' to issue %v.\n", reviewLabelName, *issue.Number)
		}

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			pr = scm.CreatePullRequest(issue, LinkType)
		}

		scm.AssignPullRequestReviewer(pr, reviewer)

		chat := chat.GetClient()

		if chat != nil {
			chat.DirectMessage(
				it.GetUserRealName(reviewer),
				fmt.Sprintf("%v has requested your review on %v", it.GetCurrentUser(), pr.URL),
			)
		}

		if util.Confirm("Open pull request in browser?") {
			util.Openbrowser(pr.URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(codereviewCmd)
	
	codereviewCmd.Flags().StringVarP(
		&LinkType,
		"link-type",
		"l",
		"close",
		"The type of link to create with the associated issue.",
	)

}
