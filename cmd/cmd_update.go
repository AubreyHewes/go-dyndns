package cmd

import (
	"github.com/AubreyHewes/update-dynamic-host/v1/config/env"
	"github.com/AubreyHewes/update-dynamic-host/v1/dynaddress"
	"github.com/AubreyHewes/update-dynamic-host/v1/log"
	dnsProviderFactory "github.com/AubreyHewes/update-dynamic-host/v1/providers/dns"
	ipProviderFactory "github.com/AubreyHewes/update-dynamic-host/v1/providers/ip"
	"github.com/urfave/cli/v2"
)

func createUpdate() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Shows additional help for the '--dns' global option",
		Action: update,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "Dynamic host to add your current IP to. i.e. great.host.me",
			},
		},
	}
}

func update(ctx *cli.Context) error {
	host := ctx.String("host")

	dynAddress, err := dynaddress.ParseHost(host)
	if err != nil {
		log.Fatalf("could not parse host: %v", err)
	}

	ipProvider, err := ipProviderFactory.NewIPProviderByName("ifconfig")
	if err != nil {
		log.Fatalf("ip provider error: %v", err)
	}

	ip, err := ipProvider.GetIP()
	if err != nil {
		log.Fatalf("could not get IP: %v", err)
	}
	log.Infof("Got IP: %s", ip)

	dynAddress.IP = ip
	dynAddress.TTL = int(env.GetOrDefaultInt("DYNAMIC_HOST_TTL", 60))

	log.Infof("Got dynAddress: %v", dynAddress)

	dnsProvider, err := dnsProviderFactory.NewDNSProviderByName("transip")
	if err != nil {
		log.Fatalf("dns provider error: %v", err)
	}

	err = dnsProvider.Update(dynAddress)
	if err != nil {
		log.Fatalf("dns provider error: %v", err)
	}

	log.Infof("It has been done!")
	return nil
}
