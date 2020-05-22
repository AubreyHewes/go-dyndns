package cmd

import (
	//"github.com/AubreyHewes/ledgo/v1/ledgo"
	"github.com/urfave/cli/v2"
)

func CreateFlags(defaultPath string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "config-dir",
			Usage: "Directory to use for storing configuration.",
			Value: defaultPath,
		},
		&cli.BoolFlag{
			Name:  "dry",
			Usage: "Dry run.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "v",
			Usage: "Verbose (add more for more verbosity)",
			Value: false,
		},
	}
}
