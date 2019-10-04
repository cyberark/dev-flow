package service

import (
	"os/exec"
)

type Bash struct{}

func (bash Bash) RunCommand(cmd string) (string, error) {
	stdout, err := exec.Command("bash", "-c", cmd).Output()
	output := string(stdout)

	if err != nil {
		return output, err
	}

	return output, nil
}
