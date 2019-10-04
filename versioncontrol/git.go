package versioncontrol

import (
	"fmt"
	"strings"

	"github.com/cyberark/dev-flow/common"
	"github.com/cyberark/dev-flow/service"
)

type Git struct{
	BashService service.Bash
}

func (git Git) Repo() (common.Repo, error) {
	cmd := "git remote show origin -n | grep h.URL | sed 's/.*://;s/.git$//'"
	
	repoSlug, err := git.BashService.RunCommand(cmd)

	if err != nil {
		return common.Repo{}, err
	}
	
	repoSlugSplit := strings.Split(strings.TrimSpace(repoSlug), "/")

	repo := common.Repo {
		Owner: repoSlugSplit[0],
		Name: repoSlugSplit[1],
	}
	
	return repo, nil
}

func (git Git) CurrentBranch() (string, error) {
	cmd := "git branch | grep \\* | cut -d ' ' -f2"
	
	branch, err := git.BashService.RunCommand(cmd)

	if err != nil {
		return branch, err
	}
	
	return strings.TrimSpace(branch), nil
}

func (git Git) Pull() (string, error) {
	return git.BashService.RunCommand("git pull")
}

func (git Git) CheckoutAndPull(branchName string) (string, error) {
	cmd := fmt.Sprintf("git checkout %v", branchName)
	checkoutOutput, err := git.BashService.RunCommand(cmd)

	output := ""
	output += checkoutOutput
	
	if err != nil {
		return output, err
	}
	
	pullOutput, err := git.Pull()

	output += pullOutput

	if err != nil {
		return output, err
	}

	return output, nil
}

func (git Git) IsRemoteBranch(branchName string) (bool, error) {
	repo, err := git.Repo()

	if err != nil {
		return false, err
	}
	
	repoSlug := fmt.Sprintf("%v/%v", repo.Owner, repo.Name)
	
	cmd := fmt.Sprintf("git ls-remote --heads git@github.com:%v.git %v", repoSlug, branchName)
	
	output, err := git.BashService.RunCommand(cmd)

	if err != nil {
		return false, err
	}

	return output != "", nil
}

func (git Git) InitBranch(issueNum int, branchName string) (string, error) {
	cmd := fmt.Sprintf("git checkout -b %v", branchName)
	checkoutOutput, err := git.BashService.RunCommand(cmd)

	output := ""
	output += checkoutOutput

	if err != nil {
		return output, err
	}
	
	cmd = fmt.Sprintf("git push --set-upstream origin %v", branchName)
	pushOutput, err := git.BashService.RunCommand(cmd)

	output += pushOutput
	
	if err != nil {
		return output, err
	}

	return output, nil
}

func (git Git) DeleteRemoteBranch(branchName string) (string, error) {
	cmd := fmt.Sprintf("git push origin --delete %v", branchName)
	return git.BashService.RunCommand(cmd)
}

func (git Git) DeleteLocalBranch(branchName string) (string, error) {
	cmd := fmt.Sprintf("git branch -D %v", branchName)
	return git.BashService.RunCommand(cmd)
}
