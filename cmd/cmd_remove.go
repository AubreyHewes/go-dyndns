package cmd

import (
	"github.com/AubreyHewes/update-dynamic-host/v1/dynaddress"
	"github.com/AubreyHewes/update-dynamic-host/v1/log"
	dns2 "github.com/AubreyHewes/update-dynamic-host/v1/providers/dns"
	"github.com/urfave/cli/v2"
)

func createRemove() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Shows additional help for the '--dns' global option",
		Action: remove,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "Host to add your ip to.",
			},
		},
	}
}

func remove(ctx *cli.Context) error {
	host := ctx.String("host")

	dynAddress, err := dynaddress.ParseHost(host)
	if err != nil {
		log.Fatalf("could not parse host: %v", err)
	}

	provider, err := dns2.NewDNSProviderByName("transip")
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = provider.Remove(dynAddress)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return nil
}
