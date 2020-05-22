// CLI application for changing logitech device led's
package main

import (
	"fmt"
	"github.com/AubreyHewes/update-dynamic-host/v1/cmd"
	"log"
	"os"
	"path/filepath"

	//"path/filepath"
	"runtime"

	//"github.com/AubreyHewes/ledgo/v1/cmd"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
)

func main() {
	app := &cli.App{}
	app.Name = "update-dynamic-host"
	app.HelpName = "update-dynamic-host"
	app.Usage = "Dynamic IP controller"
	app.EnableBashCompletion = true

	app.Version = version
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("update-dynamic-host version %s %s/%s\n", c.App.Version, runtime.GOOS, runtime.GOARCH)
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	defaultPath := ""
	cwd, err := os.Getwd()
	if err == nil {
		defaultPath = filepath.Join(cwd, ".update-dynamic-host")
	}
	app.Flags = cmd.CreateFlags(defaultPath)

	//app.Before = cmd.Before

	app.Commands = cmd.CreateCommands()

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
