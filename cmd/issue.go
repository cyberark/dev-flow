package cmd

import (
	"log"
	
	"github.com/spf13/cobra"

	"github.com/cyberark/dev-flow/issuetracking"
	"github.com/cyberark/dev-flow/util"
	"github.com/cyberark/dev-flow/versioncontrol"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Open the specified issue.",
	Long:  "Open the specified issue.",
	Args:  cobra.MinimumNArgs(1),
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

		util.Openbrowser(issue.URL)
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
}
