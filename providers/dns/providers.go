package dns

import (
	"fmt"
	"github.com/AubreyHewes/update-dynamic-host/v1/dynaddress"
	"github.com/AubreyHewes/update-dynamic-host/v1/providers/dns/transip"
)

type Provider interface {
	Update(dynAddress *dynaddress.DynAddress) error
	Remove(dynAddress *dynaddress.DynAddress) error
}

// NewDNSProviderByName Factory for DNS providers
func NewDNSProviderByName(name string) (Provider, error) {
	switch name {

	case "transip":
		return transip.NewDNSProvider()

	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
