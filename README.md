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
brew tap
brew install
```

**With Go**
```
go get
```

**Manually**
Download from releases


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

