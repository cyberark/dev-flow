# dev-flow

The dev-flow CLI is a tool for standardizing and automating common development
tasks. It currently only supports GitHub for managing issues and pull requests,
but is built to be easily extendible for additional tooling.

[![GitHub release](https://img.shields.io/github/release/cyberark/dev-flow.svg)](https://github.com/cyberark/dev-flow/releases/latest)

[![pipeline status](https://gitlab.com/cyberark/dev-flow/badges/master/pipeline.svg)](https://gitlab.com/cyberark/dev-flow/pipelines)
[![Maintainability](https://api.codeclimate.com/v1/badges/2fbe5ba2a5ac283854f0/maintainability)](https://codeclimate.com/github/cyberark/dev-flow/maintainability)

---

## Setup

### Install Golang

If you haven't already, follow the Go [installation instructions](https://golang.org/doc/install#install).

### Install dev-flow

Install `dev-flow` like so:

```
go get github.com/cyberark/dev-flow
cd $GOPATH/src/github.com/cyberark/dev-flow
go install
```

### Provide a GitHub Access Token

`dev-flow` makes heavy use of GitHub and requires that a GitHub access token be
provided in the `GITHUB_ACCESS_TOKEN` environment variable. The following setup
describes one way to provide this token securely using the OSX keychain.

1. Create a [GitHub access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)
if you haven't already.

1. Install [Summon](https://github.com/cyberark/summon) and the [summon-keyring](https://github.com/conjurinc/summon-keyring) provider.

1. Store the GitHub access token in your OSX keychain:

    ```
    $ security add-generic-password -s "summon" -a "github/access_token" -w "insert-token-here"
    ```

1. Create `~/.df-secrets.yml` to store a reference to your token:

    ```
    GITHUB_ACCESS_TOKEN: !var github/access_token
    ```

1. Create an alias to run `dev-flow` with Summon:

    ```
    alias df='summon -p keyring.py -f ~/.df-secrets.yml dev-flow'
    ```

That's it! You should now be able to use that alias to run `dev-flow` with the
secrets it needs.

### Provide a Slack API Token

`dev-flow` can be configured to deliver notifications via Slack bot to users
involved with an issue when the state of an issue changes. To enable these
notifications, you must provide the API token for a bot in the `SLACK_API_TOKEN`
environment variable.

1. Obtain the token for your Slack org's `dev-flow` app or, if need be, [create an app](https://api.slack.com/slack-apps)
yourself and retrieve its API token.

1. Store the Slack API token in your OSX keychain:

    ```
    $ security add-generic-password -s "summon" -a "slack/api_token" -w "insert-token-here"
    ```

1. Add the API token to `~/.df-secrets.yml`:

    ```
    SLACK_API_TOKEN: !var slack/api_token
    ```

`dev-flow` should now be able to send messages to users when their attention is
needed on an issue.

### Configure Labels

`dev-flow` can apply labels during the lifecycle of a story. You can provide the
names of these labels by creating `~/.df-config.yml` like so:

```
labels:
  in_progress: 'in progress'
  in_review: 'review'
```

You must create these labels in your issue tracker before using them as
`dev-flow` will not create them automatically.

## Usage

Once `dev-flow` is installed, the following commands can be run from the root directory of a source-controlled project:

- `issues`: list open issues.
- `issue [issue-key]`: open issue in browser.
- `start [issue-num]`: create branch, assign issue to self and update labels.
- `pullrequest` (`pr`): create pull request for current branch into `master`.
- `codereview [username]` (`cr`): create pull request into `master` and assign issue to user.
- `revise`: reject pull request and assign issue back to pull request creator.
- `complete`: merge pull request and (optionally) delete remote and local branches.

## Sample Workflow

Alice and Bob both work on the CoolProject team at CoolOrg. They recently
installed and configured `dev-flow` to automate some of the repetitive tasks
that they must perform on a daily basis when contributing to CoolProject.

Alice just finished wrapping up her most recent task and decides it's time to
find her next one. She takes a look at the current issues in the CoolProject
repository:

```
$ df issues
52 - DRY up all the things (unassigned) [needs info]
67 - something or other needs tests (bob) [feature, in progress]
45 - fix this crazy bug! (unassigned) [bug, ready]
```

"That last one sounds like a fun challenge", Alice thinks to herself. Let's get
more detail. She runs the `issue` command to open the issue in her browser:

```
$ df issue 45
```

After reading over the issue, She rolls up her sleeves and begins working on the issue:

```
$ df start 45
Assigned issue 45 to user alice.
Added label 'in progress' to issue 45.
...
[45--fix-this-crazy-bug b6a3fba] Issue 45 Started.
Branch '45--fix-this-crazy-bug' set up to track remote branch '45--fix-this-crazy-bug' from 'origin'.
Issue started! You are now working in branch: 45--fix-this-crazy-bug
```

Just like that, she has a local branch with an automatically generated name set
up to track a remote branch. Not only that but the issue has been labeled to
play nicely with Waffle. How convenient!

She proceeds to fix the crazy bug, commiting her work when necessary. Finally
it's time to have someone review her work. She knows Bob is involved with the
current project so she creates a pull request for him to review:

```
$ df cr bob
```

Meanwhile, Bob is sitting at his desk typing away when he suddenly receives a
Slack notification from his friendly neighborhood `dev-flow` bot:

```
alice has requested your review on https://github.com/coolorg/coolproject/pull/97
```

Bob has a few minutes to spend checking out the pull request, so he opens the
handy link in the Slack message and reviews Alice's work. His review includes
a few suggested changes, so he kicks it back her way:

```
$ git checkout 45--fix-this-crazy-bug
$ df revise
```

Now it's Alice's turn to get a visit from `dev-flow` bot:

```
bob has requested changes on https://github.com/coolorg/coolproject/pull/97
```

She opens the link to read Bob's feedback and takes a few minutes to make the
requested changes. Afterwards, she passes the story back to Bob:

```
$ df cr bob
```

Bob once again receives a notification from `dev-flow` bot and opens the link in
his browser to verify the requested changes. Satisfied, he approves and merges
the story into `master`:

```
$ df complete
Are you sure you want to merge 45--fix-this-crazy-bug [y/n]: y
Merged 45--fix-this-crazy-bug into master
...
Delete remote branch 45--fix-this-crazy-bug [y/n]: y
Remote branch deleted.
...
Delete local branch 45--fix-this-crazy-bug [y/n]: y
Deleted branch 45--fix-this-crazy-bug (was 2f3579e).
Local branch deleted.
```

With Alice's story merged and his own working environment nice and clean, Bob
can continue on his merry way. Meanwhile, Alice receives one last notification
from `dev-flow` bot to let her know that her story has been merged:

```
bob has merged your pull request https://github.com/coolorg/coolproject/pull/97
```

"Thanks, dev-flow bot!", thinks Alice, before she continues with her day.

### Contributing

We welcome contributions of all kinds to this repository. For instructions on
how to get started and descriptions of our development workflows, please see our
[contributing guide](CONTRIBUTING.md).

## License

This repository is licensed under Apache License 2.0 - see [`LICENSE`](LICENSE) for more details.
