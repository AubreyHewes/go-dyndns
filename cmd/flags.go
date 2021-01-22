package cmd

import (
	"github.com/urfave/cli/v2"
)

func CreateFlags(appConfigDir string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config-dir",
			Aliases: []string{"c"},
			Usage:   "Directory to use for storing configuration.",
			Value:   appConfigDir,
		},
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "Dry run.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbose (add more for more verbosity)",
			Value:   false,
		},
	}
}
