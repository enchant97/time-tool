package cli

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func Entrypoint(appVersion string) error {
	app := &cli.Command{
		Version:               appVersion,
		Copyright:             "Copyright (c) 2026 Leo Spratt",
		Usage:                 "A simple CLI tool to get the time",
		EnableShellCompletion: true,
		DefaultCommand:        "now",
		Commands: []*cli.Command{
			{
				Name:  "now",
				Usage: "output the current time",
				Action: func(ctx context.Context, c *cli.Command) error {
					return commandNow()
				},
			},
		},
	}
	return app.Run(context.Background(), os.Args)
}
