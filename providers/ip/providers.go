package ip

import (
	"fmt"
	"net"

	"github.com/AubreyHewes/update-dynamic-host/v1/providers/ip/ifconfig"
)

type Provider interface {
	GetIP() (net.IP, error)
}

// NewIPProviderByName Factory for IP providers
func NewIPProviderByName(name string) (Provider, error) {
	switch name {

	case "ifconfig":
		return ifconfig.NewIPProvider()

	default:
		return nil, fmt.Errorf("unrecognized IP provider: %s", name)
	}
}
