package cmd

import (
	"github.com/urfave/cli/v2"
)

// CreateCommands Creates all CLI commands
func CreateCommands() []*cli.Command {
	return []*cli.Command{
		// we used to do `createX(),` but it must be a reference, but we want one liners
		// so the following is a hacky solution https://stackoverflow.com/a/30751102/7475870
		&[]cli.Command{createUpdate()}[0],
		&[]cli.Command{createRemove()}[0],
		//&[]cli.Command{createBreathe()}[0],
		//&[]cli.Command{createIntro()}[0],
	}
}
