package ssh

import (
	"context"
	"fmt"

	"github.com/blkr0se/ginger/internal/cmd"
	"github.com/urfave/cli/v3"
)

func SshCmd() *cli.Command {
	return &cli.Command{
		Name:  "ssh",
		Usage: "connect to a server via ssh",
		Action: func(ctx context.Context, c *cli.Command) error {
			conn := cmd.NewConnector(c)
			connString := fmt.Sprintf("%s@%s", conn.User, conn.Ip)

			if err := conn.Interactive(connString); err != nil {
				return err
			}

			return nil
		},
	}
}
