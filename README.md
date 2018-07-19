# dev-flow

The dev-flow CLI is a tool for standardizing and automating common development tasks.

## Setup

If you haven't already, follow the Go [installation instructions](https://golang.org/doc/install#install).

Install `dev-flow` like so:

```
go get github.com/conjurinc/dev-flow
cd $GOPATH/src/github.com/conjurinc/dev-flow
go install
```

If you don't have one already, you will need a [GitHub access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) for your account to interact with GitHub via command line.

If you wish to notify users when their attention is required on a pull request, you can also provide an API token for a [Slack bot](https://my.slack.com/apps/A0F7YS25R-bots).

Once you have obtained these tokens, place them in a `~/.dev-flow` file:

```
github:
  access_token: [github-access-token]
  
slack:
  api_token: [slack-bot-api-token]
```

## Usage

Once `dev-flow` is installed, the following commands can be run from the root directory of a source-controlled project:

`wf issues`: list open issues.
`wf start`: create a branch for an issue, perform an initial commit, and assign the issue to the current user.
`wf pr`: create a pull request for the current branch into `master`.
`wf cr [username]`: create a pull request for the current branch into `master` and assign the associated issue to a specified user.
`wf revise`: reject a pull request and assign the associated issue back to the pull request creator.
`wf complete`: merge pull request and (optionally) delete the remote and local branches.

## Sample Workflow

Coming soon...

### Contributing

1. Fork it
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Commit your changes (`git commit -am 'Added some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new Pull Request

## License

Copyright 2018 CyberArk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
