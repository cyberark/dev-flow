package versioncontrol

import (
	"fmt"

	"os/exec"
	"strings"
)

type Git struct{}

func (git Git) runCommand(cmd string, print bool) string {
	output, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		panic(fmt.Sprintf("Failed to execute command: %s", cmd))
	}

	outputStr := string(output)

	if print {
		fmt.Println(outputStr)
	}

	return outputStr
}

func (git Git) Repo() Repo {
	slug := git.runCommand("git remote show origin -n | grep h.URL | sed 's/.*://;s/.git$//'", false)
	slugSplit := strings.Split(strings.TrimSpace(slug), "/")

	repo := Repo {
		Owner: slugSplit[0],
		Name: slugSplit[1],
	}
	
	return repo
}

func (git Git) CurrentBranch() string {
	output := git.runCommand("git branch | grep \\* | cut -d ' ' -f2", false)
	return strings.TrimSpace(output)
}

func (git Git) Pull() {
	git.runCommand("git pull", true)
}

func (git Git) CheckoutAndPull(branchName string) {
	cmd := fmt.Sprintf("git checkout %v", branchName)
	git.runCommand(cmd, true)
	git.Pull()
}

func (git Git) IsRemoteBranch(branchName string) bool {
	repo := git.Repo()
	slug := fmt.Sprintf("%v/%v", repo.Owner, repo.Name)
	
	cmd := fmt.Sprintf("git ls-remote --heads git@github.com:%v.git %v", slug, branchName)
	
	output := git.runCommand(cmd, false)

	return output != ""
}

func (git Git) InitBranch(issueNum int, branchName string) {
	cmd := fmt.Sprintf("git checkout -b %v", branchName)
	git.runCommand(cmd, true)
	
	cmd = fmt.Sprintf("git push --set-upstream origin %v", branchName)
	git.runCommand(cmd, true)
}

func (git Git) DeleteRemoteBranch(branchName string) {
	cmd := fmt.Sprintf("git push origin --delete %v", branchName)
	git.runCommand(cmd, true)
}

func (git Git) DeleteLocalBranch(branchName string) {
	cmd := fmt.Sprintf("git branch -D %v", branchName)
	git.runCommand(cmd, true)
}
