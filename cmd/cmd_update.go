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

	dnsRecord, err := dynaddress.ParseHost(host)
	if err != nil {
		log.Fatalf("could not parse host: %v", err)
	}

	dnsProvider, err := dnsProviderFactory.NewDNSProviderByName("transip")
	if err != nil {
		log.Fatalf("DNS Provider error: %v", err)
	}

	ipProvider, err := ipProviderFactory.NewIPProviderByName("default")
	if err != nil {
		log.Fatalf("IP Provider error: %v", err)
	}

	ips, err := ipProvider.GetIPs()
	if err != nil {
		log.Fatalf("Could not get IPs: %v", err)
	}
	log.Infof("Got IPs: %s", ips)

	for i := range ips {
		ip := ips[i]
		dnsRecord.IP = ip.IP.String()
		dnsRecord.TTL = int(env.GetOrDefaultInt("DYNAMIC_HOST_TTL", 60))
		dnsRecord.Type = ip.Type

		log.Verbosef("DNS record: %v", dnsRecord)

		err = dnsProvider.Update(dnsRecord)
		if err != nil {
			log.Fatalf("DNS Provider error: %v", err)
		}
	}

	log.Infof("It has been done!")
	return nil
}
