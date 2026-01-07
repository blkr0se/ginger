package main

import (
	"context"
	"log"
	"os"

	"github.com/blkr0se/ginger/cmd/flags"
	"github.com/blkr0se/ginger/cmd/ssh"
	"github.com/blkr0se/ginger/cmd/upload"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "ginger",
		// UseShortOptionHandling,
		Usage: "manage a remote server/homelab",
		Flags: []cli.Flag{
			flags.ServerUsernameFlag,
			flags.ServerHostIpFlag,
			flags.ServerSshPortFlag,
			flags.ServerSshKeyFlag,
		},
		Commands: []*cli.Command{
			upload.UploadCmd(),
			ssh.SshCmd(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
