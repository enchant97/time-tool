package cli

import (
	"context"
	"errors"
	"os"

	"github.com/enchant97/time-tool/core"
	"github.com/enchant97/time-tool/tui"
	"github.com/urfave/cli/v3"
)

func Entrypoint(appVersion string) error {
	config, err := core.ReadConfig()
	if errors.Is(err, os.ErrNotExist) {
		config.DefaultUnset()
		core.WriteConfig(config)
	} else if err != nil {
		return err
	}
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
						Name: "layout",
					},
					&cli.StringFlag{
						Name:  "location",
						Usage: "IANA Time Zone Location name",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					timeLayout := core.DefaultIfUnset(c.String("layout"), config.Layout, "")
					timeLocation := core.DefaultIfUnset(c.String("location"), config.Location, "")
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
