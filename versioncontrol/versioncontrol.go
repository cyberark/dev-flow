package versioncontrol

import (
	"github.com/cyberark/dev-flow/common"
)

type VersionControlClient interface {
	Repo() (common.Repo, error)
	CurrentBranch() (string, error)
	Pull() (string, error)
	CheckoutAndPull(string) (string, error)
	IsRemoteBranch(string) (bool, error)
	InitBranch(int, string) (string, error)
	DeleteRemoteBranch(string) (string, error)
	DeleteLocalBranch(string) (string, error)
}

func GetClient() VersionControlClient {
	return Git{}
}

type Branch struct {
	Name *string
}
