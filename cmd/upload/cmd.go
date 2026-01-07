package upload

import (
	"context"

	"github.com/blkr0se/ginger/internal/cmd"
	"github.com/blkr0se/ginger/internal/upload"
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
			conn := cmd.NewConnector(c)
			svc, err := upload.NewUploadService(conn)
			if err != nil {
				return err
			}

			req := upload.NewUploadRequest(c.String("src"), c.String("dst"))

			if err := svc.Upload(ctx, req); err != nil {
				return err
			}

			return nil
		},
	}
}
