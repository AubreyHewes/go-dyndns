package dynaddress

import (
	"golang.org/x/net/publicsuffix"
	"net"
	"strings"
)

type DynAddress struct {
	// The requested host
	Host string `json:"host"`
	// The domain of the requested host
	Domain string `json:"domain"`
	// The subdomain of the requested host or possibly empty if the requested host is the domain
	// i.e. test.me would give an empty subdomain so you would have to add the record on the apex (i.e. as "@")
	SubDomain string `json:"sub-domain"`
	// The TTL for the DNS record.. default is 300s (5mins).
	// It could be that your DNS provider does not support such a low TTL
	TTL int    `json:"ttl"`
	IP  net.IP `json:"ip"`
}

func ParseHost(host string) (*DynAddress, error) {
	host = strings.TrimSpace(host)

	dynAddress := DynAddress{}

	//etld+1
	etld1, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return nil, err
	}

	dynAddress.Host = host
	dynAddress.Domain = etld1
	dynAddress.SubDomain = strings.Replace(host, "."+etld1, "", 1)

	return &dynAddress, nil
}
