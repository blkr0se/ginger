# ginger
Ginger is a CLI for interacting with remote servers via SSH.

# Installing
`go install github.com/blkr0se/ginger`

Then test it with: `ginger -h`

## Configuring the environment
Set required environment variables:
```shell
export GINGER_SERVER_USERNAME="user" # Required
export GINGER_SERVER_HOST="192.168.1.2" # Required
export GINGER_SERVER_SSH_PORT="22" # Optional (default: 22)
export GINGER_SERVER_SSH_KEY="$HOME/.ssh/homelab-bianca" # Required
```

## Features
### Uploading a file to a remote server
`ginger upload --src hello.txt --dst hello.txt`

The `--dst` flag is optional. If not passed, the path will be mirrored on the if it already exists there.
