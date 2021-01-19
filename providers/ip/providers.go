package ip

import (
	"fmt"
	"github.com/AubreyHewes/update-dynamic-host/v1/providers/ip/default"
	"github.com/AubreyHewes/update-dynamic-host/v1/providers/ip/types"
)

// NewIPProviderByName Factory for IP providers
func NewIPProviderByName(name string) (types.Provider, error) {
	switch name {

	case "default":
		return providers_ip_default.NewIPProvider()

	default:
		return nil, fmt.Errorf("unrecognized IP provider: %s", name)
	}
}
