package upload

import (
	"context"

	"github.com/blkr0se/ginger/cmd/flags"
	"github.com/blkr0se/shun"
	"github.com/urfave/cli/v3"
)

func UploadCmd() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "upload a file to a remote server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Usage:    "path/to/file",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "dst",
				Usage: "path/to/file/in/server",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			// TODO: replace newConnector with NativeSshProvider.
			conn := newConnector(c)
			svc, err := NewUploadService(conn)
			if err != nil {
				return err
			}

			req := NewUploadRequest(c.String("src"), c.String("dst"))

			if err := svc.Upload(ctx, req); err != nil {
				return err
			}

			return nil
		},
	}
}

// Deprecated
func newConnector(c *cli.Command) *shun.SshConnector {
	return &shun.SshConnector{
		User:    c.String(flags.ServerUsernameFlag.Name),
		Ip:      c.String(flags.ServerHostIpFlag.Name),
		Port:    c.String(flags.ServerSshPortFlag.Name),
		KeyFile: c.String(flags.ServerSshKeyFlag.Name),
	}
}
