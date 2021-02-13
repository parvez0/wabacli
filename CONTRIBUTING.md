# How to Contribute

We'd love your help!

We accept contributions via GitHub pull requests, this document outlines some prerequisites which can help you get started quickly.

### Prerequisite

* Make sure you have `go 1.11+` installed with go modules activated. (You can set env var with `export GO111MODULE=on` to activate)

* Clone the source code
   ```sh
   $ git clone https://github.com/parvez0/wabacli.git
   ```
* You will also need a working whatsapp business account, you can get a new account by [following these steps](https://developers.facebook.com/docs/whatsapp/getting-started) 

Now you can build and run wabacli

### Build and run wabacli locally
```sh
├── assets
│   └── test.jpeg
├── cmd
│   └── main.go
├── config
│   ├── structure.go
│   └── viper.go
├── go.mod
├── go.sum
├── log
│   └── logger.go
└── pkg
    ├── cmd
    │   ├── context
    │   │   ├── add
    │   │   │   ├── add.go
    │   │   │   └── add_options.go
    │   │   ├── change
    │   │   │   ├── switch.go
    │   │   │   └── switch_options.go
    │   │   ├── context.go
    │   │   ├── del
    │   │   │   ├── delete.go
    │   │   │   └── delete_options.go
    │   │   ├── get
    │   │   │   ├── get.go
    │   │   │   └── get_options.go
    │   │   └── refresh
    │   │       ├── refresh.go
    │   │       └── refresh_options.go
    │   ├── root.go
    │   └── send
    │       ├── about.go
    │       ├── send.go
    │       └── send_options.go
    ├── errutil
    │   ├── badrequest
    │   │   └── validator_badrequest.go
    │   └── handler
    │       └── fatal_error.go
    ├── internal
    │   ├── handler
    │   │   ├── response_logger.go
    │   │   └── responses.go
    │   └── request
    │       ├── helper.go
    │       └── helper_test.go
    ├── tests
    │   └── *_test.go
    └── utils
        ├── helpers
        │   ├── login.go
        ├── templates
        │   ├── markdown.go
        ├── types
        │   ├── global.go
        └── validator
            └── validator.go

```
* All the commands are developed as modules, main file exits in cmd/main.go as shown in the directory structure.

* To run project cd into the cmd directory and build main.go file using the below command
```sh
    $ go build cmd/main.go -o wabacli 
```
or you can use make command to build it for you 
```sh
    $ make build
```
## Making A Change

* Consider [opening an issue](https://github.com/parvez0/wabacli/issues) before making any significant change, discussing your proposed changes ahead of time will make the contribution process smooth for everyone.

* Once the requirements are finalized you can create a new branch start developing.

* Make sure your pull request has [good commit messages](https://chris.beams.io/posts/git-commit/):
    * Separate subject from body with a blank line
    * Limit the subject line to 50 characters
    * Capitalize the subject line
    * Do not end the subject line with a period
    * Use the imperative mood in the subject line
    * Wrap the body at 72 characters
    * Use the body to explain _what_ and _why_ instead of _how_

* Try to squash unimportant commits and rebase your changes on to developed branch, this will make sure we have clean log of changes.