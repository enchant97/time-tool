package cli

import (
	"context"
	"os"

	"github.com/enchant97/time-tool/tui"
	"github.com/urfave/cli/v3"
)

func Entrypoint(appVersion string) error {
	app := &cli.Command{
		Version:               appVersion,
		Copyright:             "Copyright (c) 2026 Leo Spratt",
		Usage:                 "A simple CLI tool to get the time",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "now",
				Usage: "output the current time",
				Action: func(ctx context.Context, c *cli.Command) error {
					return commandNow()
				},
			},
			{
				Name:  "ticker",
				Usage: "output the current time at an interval",
				Action: func(ctx context.Context, c *cli.Command) error {
					return commandTicker()
				},
			},
			{
				Name:  "tui",
				Usage: "launch the terminal user interface",
				Action: func(ctx context.Context, c *cli.Command) error {
					return tui.Entrypoint()
				},
			},
		},
	}
	return app.Run(context.Background(), os.Args)
}
