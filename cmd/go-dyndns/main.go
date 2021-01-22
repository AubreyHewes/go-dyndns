// CLI application for dynamic dns
package main

import (
	"fmt"
	"github.com/AubreyHewes/go-dyndns/v1/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	name    = "go-dyndns"
	version = "dev"
)

func main() {
	app := &cli.App{}
	app.Name = name
	app.HelpName = name
	app.Usage = "Dynamic DNS CLI"
	app.EnableBashCompletion = true
	app.Version = version

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s %s/%s\n", c.App.Name, c.App.Version, runtime.GOOS, runtime.GOARCH)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	appConfigDir := fmt.Sprintf(".%s", name)

	userConfigDir, err := os.UserConfigDir()
	if err == nil {
		appConfigDir = filepath.Join(userConfigDir, appConfigDir)
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			appConfigDir = filepath.Join(cwd, appConfigDir)
		}
	}

	app.Flags = cmd.CreateFlags(appConfigDir)

	//app.Before = cmd.Before

	app.Commands = cmd.CreateCommands()

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
