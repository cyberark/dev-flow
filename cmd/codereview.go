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

var LinkTypeCodereview string = "close"

var codereviewCmd = &cobra.Command{
	Use:     "codereview [reviewer]",
	Aliases: []string{"cr"},
	Short:   "Creates a pull request and assigns a reviewer.",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		util.ValidateStringParam(
			"link-type",
			LinkTypeCodereview,
			[]string{ "close", "connect" },
		)
		
		reviewer := args[0]

		vc := versioncontrol.GetClient()
		branchName := vc.CurrentBranch()

		it := issuetracking.GetClient(vc.Repo())
		issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
		issue, err := it.GetIssue(issueKey)

		if err != nil {
			log.Fatalln(err)
		}

		err = it.RemoveIssueLabel(issue.Number, viper.GetString("labels.start"))

		if err != nil {
			log.Println(err)
		}
		
		err = it.AddIssueLabel(issue.Number, viper.GetString("labels.codereview"))

		if err != nil {
			log.Println(err)
		}

		scm := scm.GetClient()
		pr := scm.GetPullRequest(branchName)

		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			pr = scm.CreatePullRequest(*issue, LinkTypeCodereview)
		}

		scm.AssignPullRequestReviewer(pr, reviewer)

		chat := chat.GetClient()

		if chat != nil {
			login, err := it.GetCurrentUserLogin()

			if err != nil {
				log.Fatalln(err)
			}

			userRealName, err := it.GetUserRealName(reviewer)

			if err != nil {
				log.Fatalln(err)
			}
			
			chat.DirectMessage(
				userRealName,
				fmt.Sprintf("%v has requested your review on %v", login, pr.URL),
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
		&LinkTypeCodereview,
		"link-type",
		"l",
		"close",
		"The type of link to create with the associated issue.",
	)
}
