package cmd

import (
	"github.com/blkr0se/shun"
	"github.com/urfave/cli/v3"
)

var (
	ServerUsernameFlag = &cli.StringFlag{
		Name:    "username",
		Usage:   "user",
		Sources: cli.EnvVars("GINGER_SERVER_USERNAME"),
	}

	ServerHostIpFlag = &cli.StringFlag{
		Name:    "host",
		Usage:   "170.203.122.138",
		Sources: cli.EnvVars("GINGER_SERVER_HOST"),
	}

	ServerSshPortFlag = &cli.StringFlag{
		Name:    "port",
		Usage:   "1234",
		Sources: cli.EnvVars("GINGER_SERVER_SSH_PORT"),
		Value:   "22",
	}

	ServerSshKeyFlag = &cli.StringFlag{
		Name:    "key",
		Usage:   "$HOME/.ssh/your-key",
		Sources: cli.EnvVars("GINGER_SERVER_SSH_KEY"),
	}
)

func NewConnector(c *cli.Command) *shun.SshConnector {
	return &shun.SshConnector{
		User:    c.String(ServerUsernameFlag.Name),
		Ip:      c.String(ServerHostIpFlag.Name),
		Port:    c.String(ServerSshPortFlag.Name),
		KeyFile: c.String(ServerSshKeyFlag.Name),
	}
}
