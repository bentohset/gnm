Feature List
- manage ssh key
- generate ssh keys from credentials
- copy keys to remote host
- ssh into ip with keys
- connect vpn
- create macros (connect vpn + ssh)
- transfer existing keys from .ssh to hosts.yaml

Storage
- yaml config file
  - stores hosts, key filepath in .shh folder
- log file


Commands:
`gnm ls``
`gnm create`
- optional flags: a, i, d
- if alias not provided, use hostname as alias
- if keypath provided, store key dont generate key pair
- else generate key pair (ssh-keygen) and ssh-copy-id to the server
`gnm del [alias]`
`gnm use [alias]`
