# dev-flow

The dev-flow CLI is a tool for standardizing and automating common development tasks.

## Setup

### Install Golang

If you haven't already, follow the Go [installation instructions](https://golang.org/doc/install#install).

### Install dev-flow

Install `dev-flow` like so:

```
go get github.com/conjurinc/dev-flow
cd $GOPATH/src/github.com/conjurinc/dev-flow
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

## Usage

Once `dev-flow` is installed, the following commands can be run from the root directory of a source-controlled project:

- `issues`: list open issues.
- `start [issue-num]`: create branch, perform initial commit, and assign issue to self.
- `pullrequest` (`pr`): create pull request for current branch into `master`.
- `codereview [username]` (`cr`): create pull request into `master` and assign issue to user.
- `revise`: reject pull request and assign issue back to pull request creator.
- `complete`: merge pull request and (optionally) delete remote and local branches.

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
