// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import github "github.com/google/go-github/github"
import mock "github.com/stretchr/testify/mock"

import versioncontrol "github.com/cyberark/dev-flow/versioncontrol"

// GitHubService is an autogenerated mock type for the GitHubService type
type GitHubService struct {
	mock.Mock
}

// AddLabelToIssue provides a mock function with given fields: repo, issueNum, labelName
func (_m *GitHubService) AddLabelToIssue(repo versioncontrol.Repo, issueNum int, labelName string) error {
	ret := _m.Called(repo, issueNum, labelName)

	var r0 error
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, int, string) error); ok {
		r0 = rf(repo, issueNum, labelName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssignIssue provides a mock function with given fields: repo, issueNum, assigneeLogin
func (_m *GitHubService) AssignIssue(repo versioncontrol.Repo, issueNum int, assigneeLogin string) error {
	ret := _m.Called(repo, issueNum, assigneeLogin)

	var r0 error
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, int, string) error); ok {
		r0 = rf(repo, issueNum, assigneeLogin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIssue provides a mock function with given fields: repo, issueNum
func (_m *GitHubService) GetIssue(repo versioncontrol.Repo, issueNum int) (*github.Issue, error) {
	ret := _m.Called(repo, issueNum)

	var r0 *github.Issue
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, int) *github.Issue); ok {
		r0 = rf(repo, issueNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Issue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(versioncontrol.Repo, int) error); ok {
		r1 = rf(repo, issueNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIssues provides a mock function with given fields: _a0
func (_m *GitHubService) GetIssues(_a0 versioncontrol.Repo) ([]*github.Issue, error) {
	ret := _m.Called(_a0)

	var r0 []*github.Issue
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo) []*github.Issue); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.Issue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(versioncontrol.Repo) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLabel provides a mock function with given fields: repo, name
func (_m *GitHubService) GetLabel(repo versioncontrol.Repo, name string) (*github.Label, error) {
	ret := _m.Called(repo, name)

	var r0 *github.Label
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, string) *github.Label); ok {
		r0 = rf(repo, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Label)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(versioncontrol.Repo, string) error); ok {
		r1 = rf(repo, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequest provides a mock function with given fields: repo, pullRequestNum
func (_m *GitHubService) GetPullRequest(repo versioncontrol.Repo, pullRequestNum int) (*github.PullRequest, error) {
	ret := _m.Called(repo, pullRequestNum)

	var r0 *github.PullRequest
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, int) *github.PullRequest); ok {
		r0 = rf(repo, pullRequestNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.PullRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(versioncontrol.Repo, int) error); ok {
		r1 = rf(repo, pullRequestNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequests provides a mock function with given fields: repo, branchName
func (_m *GitHubService) GetPullRequests(repo versioncontrol.Repo, branchName string) ([]*github.PullRequest, error) {
	ret := _m.Called(repo, branchName)

	var r0 []*github.PullRequest
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, string) []*github.PullRequest); ok {
		r0 = rf(repo, branchName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.PullRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(versioncontrol.Repo, string) error); ok {
		r1 = rf(repo, branchName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: _a0
func (_m *GitHubService) GetUser(_a0 string) (*github.User, error) {
	ret := _m.Called(_a0)

	var r0 *github.User
	if rf, ok := ret.Get(0).(func(string) *github.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveLabelForIssue provides a mock function with given fields: repo, issueNum, labelName
func (_m *GitHubService) RemoveLabelForIssue(repo versioncontrol.Repo, issueNum int, labelName string) error {
	ret := _m.Called(repo, issueNum, labelName)

	var r0 error
	if rf, ok := ret.Get(0).(func(versioncontrol.Repo, int, string) error); ok {
		r0 = rf(repo, issueNum, labelName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
