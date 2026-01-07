package cmd

import (
	"github.com/blkr0se/shun"
	"github.com/urfave/cli/v3"
)

// Deprecated
func NewConnector(c *cli.Command) *shun.SshConnector {
	return &shun.SshConnector{
		User:    c.String(ServerUsernameFlag.Name),
		Ip:      c.String(ServerHostIpFlag.Name),
		Port:    c.String(ServerSshPortFlag.Name),
		KeyFile: c.String(ServerSshKeyFlag.Name),
	}
}
