# GNM - A Go Network Manager
GNM is a simple network manager - it can help you with managing VPN and SSH connections.

## Features
- create, list, delete ssh credentials
- generate ssh keys from credentials
- copy keys to remote host
- simple command to ssh into remote server
- open vscode ssh remote server

Pending:
- connect vpn
- create macros (connect vpn + ssh)
- transfer existing keys from .ssh to hosts.yaml


## Installation
Pre-requisites:
- OpenSSH installed


**Homebrew**
```
brew tap bentohset/tap
brew install gnm
```
or

```
brew install bentohset/tap/gnm
```

<!-- **With Go**
```
go get
``` -->

**Manually**
Download from [releases](https://github.com/bentohset/gnm/releases/latest)


## Commands:
**List all configured hosts**
```
gnm ls
gnm l
```

**Create a new host**
```
gnm create
gnm c
```

**Delete a host**
```
gnm del [alias]
gnm d [alias]
```

**SSH into a host**
```
gnm use [alias]
gnm u [alias]
```

**Open VSCode remote SSH**
- ensure vscode is installed and env path is configured
```
gnm code [alias]
gnm vsc [alias]
```


## TODO list
App:
- feature: open vscode remote ssh
- feature: ssh-copy-id, should be a separate command or in create?
- feature: connect vpn
- feature: transfer existing keys in .ssh to hosts.yaml

Devops:
- automate release tagging
  - write a script/Makefile? or explore alternatives
  - should tag the commit and push with version
  - manage versioning in the codebase (in /internal/constants, .github/workflows and .goreleaser.yaml)