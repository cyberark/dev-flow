package versioncontrol

import (
	"github.com/cyberark/dev-flow/common"
)

type VersionControlClient interface {
	Repo() common.Repo
	CurrentBranch() string
	Pull()
	CheckoutAndPull(string)
	IsRemoteBranch(string) bool
	InitBranch(int, string)
	DeleteRemoteBranch(string)
	DeleteLocalBranch(string)
}

func GetClient() VersionControlClient {
	return Git{}
}

type Branch struct {
	Name *string
}
