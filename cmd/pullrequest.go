package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/scm"
	"github.com/cyberark/dev-flow/util"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var LinkTypePullrequest string = "close"

var pullrequestCmd = &cobra.Command{
	Use:     "pullrequest",
	Aliases: []string{"pr"},
	Short:   "Creates a pull request for your branch.",
	Run: func(cmd *cobra.Command, args []string) {
		util.ValidateStringParam(
			"link-type",
			LinkTypePullrequest,
			[]string{ "close", "connect" },
		)
		
		vc := versioncontrol.GetClient()
		repo, err := vc.Repo()

		if err != nil {
			log.Fatalln(err)
		}
		
		branchName, err := vc.CurrentBranch()

		if err != nil {
			log.Fatalln(err)
		}

		scm := scm.GetClient(repo)
		pr := scm.GetPullRequest(branchName)

		if pr != nil {
			fmt.Println("Pull request already exists for branch", branchName)
		} else {
			issueKey := issuetracking.GetIssueKeyFromBranchName(branchName)
			issue, err := issuetracking.GetClient(repo).GetIssue(issueKey)

			if err != nil {
				log.Fatalln(err)
			}
			
			pr = scm.CreatePullRequest(*issue, LinkTypePullrequest)
		}

		if util.Confirm("Open pull request in browser?") {
			util.Openbrowser(pr.URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullrequestCmd)
	
	pullrequestCmd.Flags().StringVarP(
		&LinkTypePullrequest,
		"link-type",
		"l",
		"close",
		"The type of link to create with the associated issue.",
	)
}
