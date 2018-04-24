package versioncontrol

import (
	"fmt"

	"os/exec"
	"strings"
)

type Git struct { }

func (git Git) Repo() Repo {
	slug := git.RunCommand("git remote show origin -n | grep h.URL | sed 's/.*://;s/.git$//'")
	slugSplit := strings.Split(strings.TrimSpace(slug), "/")

	repo := Repo {
		Owner: slugSplit[0],
		Name: slugSplit[1],
	}
	
	return repo
}

func (git Git) CurrentBranch() string {
	output := git.RunCommand("git branch | grep \\* | cut -d ' ' -f2")
	return strings.TrimSpace(output)
}

func (git Git) Pull() {
	git.RunCommand("git pull")
}

func (git Git) CheckoutAndPull(branchName string) {
	cmd := fmt.Sprintf("git checkout %v", branchName)
	git.RunCommand(cmd)
	git.Pull()
}

func (git Git) IsRemoteBranch(branchName string) bool {
	repo := git.Repo()
	slug := fmt.Sprintf("%v/%v", repo.Owner, repo.Name)
	
	cmd := fmt.Sprintf("git ls-remote --heads git@github.com:%v.git %v", slug, branchName)
	
	output := git.RunCommand(cmd)

	return output != ""
}

func (git Git) InitBranch(issueNum int, branchName string) {
	cmd := fmt.Sprintf("git checkout -b %v", branchName)
	git.RunCommand(cmd)
	
	cmd = fmt.Sprintf("git commit -m 'Issue %v Started.' --allow-empty", issueNum)
	git.RunCommand(cmd)

	cmd = fmt.Sprintf("git push --set-upstream origin %v", branchName)
	git.RunCommand(cmd)
}

func (git Git) DeleteRemoteBranch(branchName string) {
	cmd := fmt.Sprintf("git push origin --delete %v", branchName)
	git.RunCommand(cmd)
}

func (git Git) DeleteLocalBranch(branchName string) {
	cmd := fmt.Sprintf("git branch -d %v", branchName)
	git.RunCommand(cmd)
}

func (git Git) RunCommand(cmd string) string {
	output, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		panic(fmt.Sprintf("Failed to execute command: %s", cmd))
	}

	outputStr := string(output)

	fmt.Println(outputStr)

	return outputStr
}
