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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "layout",
						Value: "RFC3339",
					},
					&cli.StringFlag{
						Name:  "location",
						Value: "Local",
						Usage: "IANA Time Zone Location name",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					timeLayout := c.String("layout")
					timeLocation := c.String("location")
					return commandNow(timeLayout, timeLocation)
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
