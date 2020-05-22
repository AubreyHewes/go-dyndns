// Package transip implements a DNS provider for updating the dynamic IP using TransIP.
package transip

import (
	"errors"
	"fmt"
	"github.com/AubreyHewes/update-dynamic-host/v1/config/env"
	"github.com/AubreyHewes/update-dynamic-host/v1/dynaddress"
	"github.com/AubreyHewes/update-dynamic-host/v1/log"
	domain2 "github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/repository"
	"sync"

	"github.com/transip/gotransip/v6"
)

// Config is used to configure the creation of the DNSProvider
type Config struct {
	AccountName    string
	PrivateKeyPath string
	TTL            int
}

// NewDefaultConfig returns a default configuration for the DNSProvider
func NewDefaultConfig() *Config {
	return &Config{
		TTL: int(env.GetOrDefaultInt("TRANSIP_TTL", 60)),
	}
}

// DNSProvider describes a provider for TransIP
type DNSProvider struct {
	config       *Config
	client       repository.Client
	dnsEntriesMu sync.Mutex
}

// NewDNSProvider returns a DNSProvider instance configured for TransIP.
// Credentials must be passed in the environment variables:
// TRANSIP_ACCOUNTNAME, TRANSIP_PRIVATEKEYPATH.
func NewDNSProvider() (*DNSProvider, error) {
	values, err := env.Get("TRANSIP_ACCOUNT_NAME", "TRANSIP_PRIVATE_KEY_PATH")
	if err != nil {
		return nil, fmt.Errorf("transip: %v", err)
	}

	config := NewDefaultConfig()
	config.AccountName = values["TRANSIP_ACCOUNT_NAME"]
	config.PrivateKeyPath = values["TRANSIP_PRIVATE_KEY_PATH"]

	return NewDNSProviderConfig(config)
}

// NewDNSProviderConfig return a DNSProvider instance configured for TransIP.
func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	if config == nil {
		return nil, errors.New("transip: the configuration of the DNS provider is nil")
	}

	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    config.AccountName,
		PrivateKeyPath: config.PrivateKeyPath,
	})
	if err != nil {
		return nil, fmt.Errorf("transip: %v", err)
	}

	return &DNSProvider{client: client, config: config}, nil
}

// Add creates an A record to fulfill the dns
func (d *DNSProvider) Update(dynAddress *dynaddress.DynAddress) error {

	repo := domain2.Repository{Client: d.client}

	dnsEntries, err := repo.GetDNSEntries(dynAddress.Domain)
	if err != nil {
		return fmt.Errorf("transip: %v", err)
	}

	dnsEntry := domain2.DNSEntry{
		Name:    dynAddress.SubDomain,
		Expire:  dynAddress.TTL, // d.config.TTL
		Type:    "A",
		Content: dynAddress.IP.String(),
	}

	for i := range dnsEntries {
		if dnsEntries[i].Name == dnsEntry.Name {
			// Found!
			log.Infof("Updating existing record for %v", dnsEntry)
			err = repo.UpdateDNSEntry(dynAddress.Domain, dnsEntry)

			if err != nil {
				return fmt.Errorf("transip: %v", err)
			}

			return nil
		}
	}

	log.Infof("Creating new record")
	err = repo.AddDNSEntry(dynAddress.Domain, dnsEntry)

	if err != nil {
		return fmt.Errorf("transip: %v", err)
	}

	return nil
}

// Remove removes the A record matching the specified parameters
func (d *DNSProvider) Remove(dynAddress *dynaddress.DynAddress) error {

	repo := domain2.Repository{Client: d.client}

	dnsEntries, err := repo.GetDNSEntries(dynAddress.Domain)
	if err != nil {
		return fmt.Errorf("transip: %v", err)
	}

	for i := range dnsEntries {
		if dnsEntries[i].Name == dynAddress.SubDomain {
			// Found!
			log.Infof("Removing existing record")
			err = repo.RemoveDNSEntry(dynAddress.Domain, dnsEntries[i])

			if err != nil {
				return fmt.Errorf("transip: %v", err)
			}

			return nil
		}
	}

	return nil
}
