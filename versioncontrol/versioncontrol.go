package versioncontrol

type VersionControl interface {
	Repo() Repo
	CurrentBranch() string
	Pull()
	CheckoutAndPull(string)
	IsRemoteBranch(string) bool
	InitBranch(int, string)
	DeleteRemoteBranch(string)
	DeleteLocalBranch(string)
}

func GetClient() VersionControl {
	return Git { }
}

type Repo struct {
	Owner string
	Name string
}

type Branch struct {
	Name *string
}
