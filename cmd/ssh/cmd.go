package ssh

import (
	"context"
	"fmt"

	"github.com/blkr0se/shun"
	"github.com/urfave/cli/v3"
)

func SshCmd() *cli.Command {
	return &cli.Command{
		Name:  "ssh",
		Usage: "connect to a server via ssh",
		Action: func(ctx context.Context, c *cli.Command) error {
			args := []string{
				fmt.Sprintf("%s@%s", c.String("username"), c.String("host")),
				"-p",
				c.String("port"),
			}

			provider := shun.NewBinProvider(args...)

			if err := provider.Connect(); err != nil {
				return err
			}

			if err := provider.Wait(); err != nil {
				return err
			}

			return nil
		},
	}
}
